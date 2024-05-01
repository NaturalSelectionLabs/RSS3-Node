package decentralized

import (
	"fmt"
	"net/url"
	"time"

	"github.com/rss3-network/node/schema/worker"
	"github.com/rss3-network/protocol-go/schema"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
)

type ActivityRequest struct {
	ID          string `param:"id"`
	ActionLimit int    `query:"action_limit"  validate:"min=1,max=20" default:"10"`
	ActionPage  int    `query:"action_page" validate:"min=1" default:"1"`
}

type AccountActivitiesRequest struct {
	Account        string               `param:"account" validate:"required"`
	Limit          int                  `query:"limit" validate:"min=1,max=100" default:"100"`
	ActionLimit    int                  `query:"action_limit" validate:"min=1,max=20" default:"10"`
	Cursor         *string              `query:"cursor"`
	SinceTimestamp *uint64              `query:"since_timestamp"`
	UntilTimestamp *uint64              `query:"until_timestamp"`
	Status         *bool                `query:"success"`
	Direction      *activityx.Direction `query:"direction"`
	Network        []network.Network    `query:"network"`
	Tag            []tag.Tag            `query:"tag"`
	Type           []schema.Type        `query:"-"`
	Platform       []string             `query:"platform"`
}

type ActivityResponse struct {
	Data *activityx.Activity `json:"data"`
	Meta *MetaTotalPages     `json:"meta"`
}

type ActivitiesResponse struct {
	Data []*activityx.Activity `json:"data"`
	Meta *MetaCursor           `json:"meta,omitempty"`
}

type MetaTotalPages struct {
	TotalPages int `json:"totalPages"`
}

type MetaCursor struct {
	Cursor string `json:"cursor"`
}

type StatisticResponse struct {
	Count      int64      `json:"count"`
	LastUpdate *time.Time `json:"last_update,omitempty"`
}

type WorkerResponse struct {
	Data []WorkerInfo `json:"data"`
}

type WorkerInfo struct {
	Network network.Network `json:"network"`
	Worker  worker.Worker   `json:"worker"`
	Status  worker.Status   `json:"status"`
}

func (h *Hub) parseParams(params url.Values, tags []tag.Tag) ([]schema.Type, error) {
	if len(tags) == 0 {
		return nil, nil
	}

	types := make([]schema.Type, 0)

	for _, typex := range params["type"] {
		var (
			value schema.Type
			err   error
		)

		for _, tag := range tags {
			value, err = schema.ParseTypeFromString(tag, typex)
			if err == nil {
				types = append(types, value)

				break
			}
		}

		if err != nil {
			return nil, fmt.Errorf("invalid type: %s", typex)
		}
	}

	return types, nil
}
