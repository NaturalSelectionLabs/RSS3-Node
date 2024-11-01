package info

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/node/monitor"
	"github.com/rss3-network/node/schema/worker"
	"github.com/rss3-network/node/schema/worker/decentralized"
	"github.com/rss3-network/node/schema/worker/federated"
	"github.com/rss3-network/node/schema/worker/rss"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"go.uber.org/zap"
)

type WorkerResponse struct {
	Data ComponentInfo `json:"data"`
}

type ComponentInfo struct {
	Decentralized []*WorkerInfo `json:"decentralized,omitempty"`
	RSS           *WorkerInfo   `json:"rss,omitempty"`
	Federated     []*WorkerInfo `json:"federated,omitempty"`
}

type WorkerInfo struct {
	WorkerID string          `json:"worker_id"`
	Worker   worker.Worker   `json:"worker"`
	Network  network.Network `json:"network"`
	Tags     []tag.Tag       `json:"tags"`
	Platform string          `json:"platform"`
	Status   worker.Status   `json:"status"`
	monitor.WorkerProgress
}

// GetWorkersStatus returns the status of all workers.
func (c *Component) GetWorkersStatus(ctx echo.Context) error {
	go c.CollectTrace(ctx.Request().Context(), ctx.Request().RequestURI, "status")

	workerCount := config.CalculateWorkerCount(c.config)
	workerInfoChan := make(chan *WorkerInfo, workerCount)

	// Handle redis + decentralized case first
	if c.redisClient != nil && len(c.config.Component.Decentralized) > 0 {
		c.fetchAllWorkerInfo(ctx, workerInfoChan)
		response := c.buildWorkerResponse(workerInfoChan)

		return ctx.JSON(http.StatusOK, response)
	}

	// Logic for RSS and Federated
	response := &WorkerResponse{
		Data: ComponentInfo{},
	}

	// Handle RSS if exists
	if c.config.Component.RSS != nil {
		m := c.config.Component.RSS
		response.Data.RSS = &WorkerInfo{
			WorkerID: m.ID,
			Network:  m.Network,
			Worker:   m.Worker,
			Tags:     rss.ToTagsMap[m.Worker.(rss.Worker)],
			Platform: rss.ToPlatformMap[m.Worker.(rss.Worker)].String(),
			Status:   worker.StatusReady,
		}
	}

	// Handle Federated if exists
	if len(c.config.Component.Federated) > 0 {
		f := c.config.Component.Federated[0]
		if f.Worker == federated.Core {
			response.Data.Federated = []*WorkerInfo{{
				WorkerID: f.ID,
				Network:  f.Network,
				Worker:   f.Worker,
				Tags:     federated.ToTagsMap[federated.Core],
				Platform: federated.ToPlatformMap[federated.Core].String(),
				Status:   worker.StatusReady,
			}}
		}
	}

	if response.Data.RSS != nil || len(response.Data.Federated) > 0 {
		return ctx.JSON(http.StatusOK, response)
	}

	return nil
}

// fetchAllWorkerInfo fetches the status of all workers concurrently.
func (c *Component) fetchAllWorkerInfo(ctx echo.Context, workerInfoChan chan<- *WorkerInfo) {
	var wg sync.WaitGroup

	fetchWorkerInfo := func(w *config.Module, fetchFunc func(context.Context, *config.Module) *WorkerInfo) {
		wg.Add(1)

		go func(module *config.Module) {
			defer wg.Done()

			workerInfoChan <- fetchFunc(ctx.Request().Context(), module)
		}(w)
	}

	modules := make([]*config.Module, 0, config.CalculateWorkerCount(c.config))

	if len(c.config.Component.Decentralized) > 0 {
		modules = append(modules, c.config.Component.Decentralized...)
	}

	if len(c.config.Component.Federated) > 0 {
		modules = append(modules, c.config.Component.Federated...)
	}

	if c.config.Component.RSS != nil {
		modules = append(modules, c.config.Component.RSS)
	}

	for _, m := range modules {
		if m.Network.Protocol() == network.RSSProtocol {
			if rssWorker := rss.GetValueByWorkerStr(m.Worker.Name()); rssWorker != 0 {
				m.Worker = rssWorker
			}
		}

		fetchWorkerInfo(m, c.fetchWorkerInfo)
	}

	go func() {
		wg.Wait()
		close(workerInfoChan)
	}()
}

// buildWorkerResponse builds the worker response from the worker info channel
func (c *Component) buildWorkerResponse(workerInfoChan <-chan *WorkerInfo) *WorkerResponse {
	response := &WorkerResponse{
		Data: ComponentInfo{
			Decentralized: []*WorkerInfo{},
			Federated:     []*WorkerInfo{},
		},
	}

	for workerInfo := range workerInfoChan {
		switch workerInfo.Network.Protocol() {
		case network.RSSProtocol:
			if c.config.Component.RSS != nil {
				response.Data.RSS = workerInfo
			}
		case network.EthereumProtocol, network.FarcasterProtocol, network.ArweaveProtocol, network.NearProtocol:
			response.Data.Decentralized = append(response.Data.Decentralized, workerInfo)
		case network.ActivityPubProtocol:
			response.Data.Federated = append(response.Data.Federated, workerInfo)
		default:
		}
	}

	return response
}

// fetchWorkerInfo fetches the worker info with the different network protocol.
func (c *Component) fetchWorkerInfo(ctx context.Context, module *config.Module) *WorkerInfo {
	if module == nil {
		zap.L().Info("params module is nil in fetchWorkerInfo")

		return &WorkerInfo{
			WorkerID: "",
			Status:   worker.StatusUnknown,
		}
	}

	// Fetch status and progress from a specific worker by id.
	status, workerProgress := c.getWorkerStatusAndProgressByID(ctx, module.ID)

	workerInfo := &WorkerInfo{
		WorkerID: module.ID,
		Network:  module.Network,
		Worker:   module.Worker,
		Status:   status,
		WorkerProgress: monitor.WorkerProgress{
			RemoteState:  workerProgress.RemoteState,
			IndexedState: workerProgress.IndexedState,
			IndexCount:   workerProgress.IndexCount,
		},
	}

	switch module.Network.Protocol() {
	case network.ActivityPubProtocol:
		if federatedWorker, ok := module.Worker.(federated.Worker); ok {
			workerInfo.Platform = federated.ToPlatformMap[federatedWorker].String()
			workerInfo.Tags = federated.ToTagsMap[federatedWorker]
		}
	case network.EthereumProtocol, network.ArweaveProtocol, network.FarcasterProtocol, network.NearProtocol:
		if decentralizedWorker, ok := module.Worker.(decentralized.Worker); ok {
			workerInfo.Platform = decentralized.ToPlatformMap[decentralizedWorker].String()
			workerInfo.Tags = decentralized.ToTagsMap[decentralizedWorker]

			// Handle special tags for decentralized core workers.
			if decentralizedWorker == decentralized.Core {
				switch module.Network {
				case network.Farcaster:
					workerInfo.Tags = []tag.Tag{tag.Social}
				case network.Arweave:
					workerInfo.Tags = []tag.Tag{tag.Transaction}
				case network.VSL:
					workerInfo.Tags = append(workerInfo.Tags, tag.Exchange)
				default:
				}
			}
		}
	case network.RSSProtocol:
		if rssWorker, ok := module.Worker.(rss.Worker); ok {
			workerInfo.Platform = rss.ToPlatformMap[rssWorker].String()
			workerInfo.Tags = rss.ToTagsMap[rssWorker]
		}
	}

	return workerInfo
}

// getWorkerStatusAndProgressByID gets both worker status and progress from Redis cache by worker ID.
func (c *Component) getWorkerStatusAndProgressByID(ctx context.Context, workerID string) (worker.Status, monitor.WorkerProgress) {
	if c.redisClient == nil {
		return worker.StatusUnknown, monitor.WorkerProgress{}
	}

	statusKey := c.buildWorkerIDStatusCacheKey(workerID)
	progressKey := c.buildWorkerProgressCacheKey(workerID)

	command := c.redisClient.B().Mget().Key(statusKey, progressKey).Build()

	result := c.redisClient.Do(ctx, command)
	if err := result.Error(); err != nil {
		return worker.StatusUnknown, monitor.WorkerProgress{}
	}

	values, err := result.ToArray()
	if err != nil || len(values) < 2 {
		return worker.StatusUnknown, monitor.WorkerProgress{}
	}

	// Parse the status
	statusValue, err := c.parseRedisJSONValue(values[0].String())
	if err != nil {
		return worker.StatusUnknown, monitor.WorkerProgress{}
	}

	status, err := worker.StatusString(statusValue)
	if err != nil {
		status = worker.StatusUnknown
	}

	// Parse the progress
	progressValue, err := c.parseRedisJSONValue(values[1].String())
	if err != nil {
		return status, monitor.WorkerProgress{}
	}

	var workerProgress monitor.WorkerProgress

	if progressValue != "" {
		err = json.Unmarshal([]byte(progressValue), &workerProgress)
		if err != nil {
			return status, monitor.WorkerProgress{}
		}
	}

	return status, workerProgress
}

// extract the value field from the redis result string
func (c *Component) parseRedisJSONValue(jsonStr string) (string, error) {
	var data map[string]interface{}

	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return "", err
	}

	value, ok := data["Value"].(string)
	if !ok {
		return "", fmt.Errorf("value field is not a string")
	}

	return value, nil
}

// buildWorkerIDStatusCacheKey builds the cache key for the worker status by id.
func (c *Component) buildWorkerIDStatusCacheKey(workerID string) string {
	return fmt.Sprintf("worker:status:id:%s", workerID)
}

// buildWorkerProgressCacheKey builds the cache key for the worker progress by id.
func (c *Component) buildWorkerProgressCacheKey(workerID string) string {
	return fmt.Sprintf("worker:progress:%s", workerID)
}
