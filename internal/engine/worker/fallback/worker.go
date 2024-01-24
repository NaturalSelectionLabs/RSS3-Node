package fallback

import (
	"fmt"

	"github.com/rss3-network/serving-node/config"
	"github.com/rss3-network/serving-node/internal/engine"
	"github.com/rss3-network/serving-node/internal/engine/worker/fallback/arweave"
	"github.com/rss3-network/serving-node/internal/engine/worker/fallback/ethereum"
	"github.com/rss3-network/serving-node/schema/filter"
)

// NewWorker creates a new fallback worker.
func NewWorker(config *config.Module) (engine.Worker, error) {
	switch config.Network.Source() {
	case filter.NetworkEthereumSource:
		return ethereum.NewWorker(config)
	case filter.NetworkArweaveSource:
		return arweave.NewWorker(config)
	default:
		return nil, fmt.Errorf("unsupported worker %s", config.Network)
	}
}
