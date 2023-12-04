package arweave

import (
	"context"
	"fmt"
	"math/big"

	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/source/arweave"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/naturalselectionlabs/rss3-node/schema/metadata"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

var _ engine.Worker = (*worker)(nil)

type worker struct {
	config *engine.Config
}

func (w *worker) Name() string {
	return engine.Fallback.String()
}

// Match returns true if the task is an Arweave task.
func (w *worker) Match(_ context.Context, task engine.Task) (bool, error) {
	return task.Network() == filter.NetworkArweave, nil
}

// Transform returns a feed with the action of the task.
func (w *worker) Transform(ctx context.Context, task engine.Task) (*schema.Feed, error) {
	// Cast the task to an Arweave task.
	arweaveTask, ok := task.(*arweave.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	// Build the feed.
	feed, err := task.BuildFeed()
	if err != nil {
		return nil, fmt.Errorf("build feed: %w", err)
	}

	// If the task is a native transfer transaction, handle it.
	if w.matchArweaveNativeTransferTransaction(arweaveTask) {
		// Handle the native transfer transaction.
		action, err := w.handleArweaveNativeTransferTransaction(ctx, arweaveTask)
		if err != nil {
			return nil, fmt.Errorf("handle native transfer transaction: %w", err)
		}

		feed.Type = action.Type
		feed.Actions = append(feed.Actions, action)
	}

	return feed, nil
}

// matchArweaveNativeTransferTransaction returns true if the transaction is a native transfer transaction.
func (w *worker) matchArweaveNativeTransferTransaction(task *arweave.Task) bool {
	// Parse the transaction quantity
	value, ok := new(big.Int).SetString(task.Transaction.Quantity, 10)
	if !ok {
		return ok
	}

	// Return true if the value is positive.
	return value.Sign() == 1
}

// handleArweaveNativeTransferTransaction returns the action of the native transfer transaction.
func (w *worker) handleArweaveNativeTransferTransaction(ctx context.Context, task *arweave.Task) (*schema.Action, error) {
	value, ok := new(big.Int).SetString(task.Transaction.Quantity, 10)
	if !ok {
		return nil, fmt.Errorf("parse transaction quantity %s", task.Transaction.Quantity)
	}

	// Build the native transfer transaction action.
	return w.buildArweaveTransactionTransferAction(ctx, task.Transaction.Owner, task.Transaction.Target, value)
}

// buildArweaveTransactionTransferAction returns the native transfer transaction action.
func (w *worker) buildArweaveTransactionTransferAction(_ context.Context, from, to string, tokenValue *big.Int) (*schema.Action, error) {
	action := schema.Action{
		Type: filter.TypeTransactionTransfer,
		From: from,
		To:   to,
		Metadata: metadata.TransactionTransfer{
			Value: lo.ToPtr(decimal.NewFromBigInt(tokenValue, 0)),
		},
	}

	return &action, nil
}

// NewWorker returns a new Arweave worker.
func NewWorker(config *engine.Config) (engine.Worker, error) {
	var instance = worker{
		config: config,
	}

	return &instance, nil
}
