package kiwistand

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/protocol/ethereum"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract"
	"github.com/rss3-network/node/provider/ethereum/contract/kiwistand"
	"github.com/rss3-network/node/provider/ethereum/token"
	"github.com/rss3-network/node/schema/worker/decentralized"
	"github.com/rss3-network/protocol-go/schema"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

// Worker is the worker for KiwiStand.
var _ engine.Worker = (*worker)(nil)

type worker struct {
	config                  *config.Module
	ethereumClient          ethereum.Client
	tokenClient             token.Client
	protocolRewardsFilterer *kiwistand.ProtocolRewardsFilterer
	kiwiFilterer            *kiwistand.KiwiFilterer
}

func (w *worker) Name() string {
	return decentralized.KiwiStand.String()
}

func (w *worker) Platform() string {
	return decentralized.PlatformKiwiStand.String()
}

func (w *worker) Network() []network.Network {
	return []network.Network{
		network.Optimism,
	}
}

func (w *worker) Tags() []tag.Tag {
	return []tag.Tag{
		tag.Collectible,
		tag.Transaction,
		tag.Social,
	}
}

func (w *worker) Types() []schema.Type {
	return []schema.Type{
		typex.CollectibleMint,
		typex.TransactionTransfer,
		typex.SocialMint,
	}
}

// Filter kiwistand contract address and event hash.
func (w *worker) Filter() engine.DataSourceFilter {
	return &source.Filter{
		LogAddresses: []common.Address{
			kiwistand.AddressKIWI,
			kiwistand.AddressProtocolRewards,
		},
		LogTopics: []common.Hash{
			kiwistand.EventHashRewardsDeposit,
			kiwistand.EventHashTransfer,
			kiwistand.EventHashSale,
		},
	}
}

// Transform Ethereum task to activityx.
func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	zap.L().Debug("transforming KiwiStand task", zap.String("task_id", task.ID()))

	ethereumTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	// Build default kiwistand activity from task.
	activity, err := ethereumTask.BuildActivity(activityx.WithActivityPlatform(w.Platform()))
	if err != nil {
		return nil, fmt.Errorf("build activity: %w", err)
	}

	// Match and handle ethereum logs.
	for _, log := range ethereumTask.Receipt.Logs {
		var (
			actions []*activityx.Action
			err     error
		)

		zap.L().Debug("processing log",
			zap.String("task_id", ethereumTask.ID()),
			zap.String("log_address", log.Address.String()))

		// Match kiwistand core contract events
		switch {
		case w.matchRewardsDeposit(ethereumTask, log):
			zap.L().Debug("processing rewards deposit event")

			actions, err = w.transformRewardsDeposit(ctx, ethereumTask, log)
		case w.matchTransfer(ethereumTask, log):
			zap.L().Debug("processing transfer event")

			actions, err = w.transformKiwiMint(ctx, ethereumTask, log)
		case w.matchSale(ethereumTask, log):
			zap.L().Debug("processing sale event")

			actions, err = w.transformSale(ctx, ethereumTask, log)
		case w.matchMintComment(ethereumTask, log):
			zap.L().Debug("processing mint comment event")

			actions, err = w.transformMintComment(ctx, ethereumTask, log)
		default:
			continue
		}

		if err != nil {
			return nil, err
		}

		activity.Type = typex.CollectibleMint

		activity.Actions = append(activity.Actions, actions...)
	}

	zap.L().Debug("successfully transformed KiwiStand task")

	return activity, nil
}

// matchRewardsDeposit matches the rewards deposit event.
func (w *worker) matchRewardsDeposit(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == kiwistand.AddressProtocolRewards && contract.MatchEventHashes(log.Topics[0], kiwistand.EventHashRewardsDeposit)
}

// matchTransfer matches the transfer event.
func (w *worker) matchTransfer(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == kiwistand.AddressKIWI && contract.MatchEventHashes(log.Topics[0], kiwistand.EventHashTransfer)
}

// matchSale matches the sale event.
func (w *worker) matchSale(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == kiwistand.AddressKIWI && contract.MatchEventHashes(log.Topics[0], kiwistand.EventHashSale)
}

// matchMintComment matches the mint comment event.
func (w *worker) matchMintComment(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == kiwistand.AddressKIWI && contract.MatchEventHashes(log.Topics[0], kiwistand.EventHashMintComment)
}

// transformKiwiMint transforms Transfer event.
func (w *worker) transformKiwiMint(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	// Parse Transfer event.
	event, err := w.kiwiFilterer.ParseTransfer(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse Transfer event: %w", err)
	}

	action, err := w.buildKiwiMintAction(ctx, task, event.From, event.To, log.Address, event.TokenId, big.NewInt(1))
	if err != nil {
		return nil, fmt.Errorf("build KiwiMint action: %w", err)
	}

	return []*activityx.Action{
		action,
	}, nil
}

// transformRewardsDeposit transforms RewardsDeposit event.
func (w *worker) transformRewardsDeposit(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.protocolRewardsFilterer.ParseRewardsDeposit(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse Transfer event: %w", err)
	}

	var actions []*activityx.Action

	var creatorRewardAction, createReferralRewardAction, mintReferralRewardAction, firstMinterRewardAction, zoraRewardAction *activityx.Action

	if event.CreatorReward.Cmp(big.NewInt(0)) > 0 {
		creatorRewardAction, err = w.buildKiwiFeeAction(ctx, task, event.From, kiwistand.AddressProtocolRewards, event.CreatorReward)
		if err != nil {
			return nil, err
		}

		actions = append(actions, creatorRewardAction)
	}

	if event.CreateReferralReward.Cmp(big.NewInt(0)) > 0 {
		createReferralRewardAction, err = w.buildKiwiFeeAction(ctx, task, event.From, kiwistand.AddressProtocolRewards, event.CreateReferralReward)
		if err != nil {
			return nil, err
		}

		actions = append(actions, createReferralRewardAction)
	}

	if event.MintReferralReward.Cmp(big.NewInt(0)) > 0 {
		mintReferralRewardAction, err = w.buildKiwiFeeAction(ctx, task, event.From, kiwistand.AddressProtocolRewards, event.MintReferralReward)
		if err != nil {
			return nil, err
		}

		actions = append(actions, mintReferralRewardAction)
	}

	if event.FirstMinterReward.Cmp(big.NewInt(0)) > 0 {
		firstMinterRewardAction, err = w.buildKiwiFeeAction(ctx, task, event.From, kiwistand.AddressProtocolRewards, event.FirstMinterReward)
		if err != nil {
			return nil, err
		}

		actions = append(actions, firstMinterRewardAction)
	}

	if event.ZoraReward.Cmp(big.NewInt(0)) > 0 {
		zoraRewardAction, err = w.buildKiwiFeeAction(ctx, task, event.From, kiwistand.AddressProtocolRewards, event.ZoraReward)

		if err != nil {
			return nil, err
		}

		actions = append(actions, zoraRewardAction)
	}

	return actions, nil
}

// transformSale transforms Sale event.
func (w *worker) transformSale(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.kiwiFilterer.ParseSale(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse sale event: %w", err)
	}

	action, err := w.buildKiwiFeeAction(ctx, task, task.Transaction.From, kiwistand.AddressKIWI, new(big.Int).Mul(event.Quantity, event.PricePerToken))
	if err != nil {
		return nil, err
	}

	return []*activityx.Action{
		action,
	}, nil
}

// transformMintComment transforms MintComment event.
func (w *worker) transformMintComment(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.kiwiFilterer.ParseMintComment(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse mint comment event: %w", err)
	}

	return []*activityx.Action{
		w.buildKiwiMintCommentAction(ctx, task.Transaction.From, kiwistand.AddressKIWI, event.Comment),
	}, nil
}

// buildKiwiMintAction builds KiwiMint action.
func (w *worker) buildKiwiMintAction(ctx context.Context, task *source.Task, from, to common.Address, contract common.Address, id *big.Int, value *big.Int) (*activityx.Action, error) {
	zap.L().Debug("building kiwi mint action",
		zap.String("from", from.String()),
		zap.String("to", to.String()),
		zap.String("contract", contract.String()),
		zap.String("token_id", id.String()),
		zap.String("value", value.String()))

	tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, &contract, id, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata: %w", err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(value, 0))

	zap.L().Debug("successfully built kiwi mint action")

	return &activityx.Action{
		Type:     typex.CollectibleMint,
		Platform: w.Platform(),
		From:     from.String(),
		To:       to.String(),
		Metadata: metadata.CollectibleTransfer(*tokenMetadata),
	}, nil
}

// buildKiwiMintCommentAction builds KiwiMintComment action.
func (w *worker) buildKiwiMintCommentAction(_ context.Context, from common.Address, to common.Address, comment string) *activityx.Action {
	zap.L().Debug("building kiwi mint comment action",
		zap.String("from", from.String()),
		zap.String("to", to.String()),
		zap.String("comment", comment))

	action := &activityx.Action{
		From:     from.String(),
		To:       lo.If(to == ethereum.AddressGenesis, "").Else(to.String()),
		Platform: w.Platform(),
		Type:     typex.SocialMint,
		Metadata: &metadata.SocialPost{
			Title: comment,
			Body:  comment,
		},
	}

	zap.L().Debug("successfully built kiwi mint comment action")

	return action
}

// buildKiwiFee builds fee
func (w *worker) buildKiwiFeeAction(ctx context.Context, task *source.Task, from common.Address, to common.Address, amount *big.Int) (*activityx.Action, error) {
	zap.L().Debug("building kiwi fee action",
		zap.String("from", from.String()),
		zap.String("to", to.String()),
		zap.String("amount", amount.String()))

	tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, nil, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata: %w", err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(amount, 0))

	zap.L().Debug("successfully built kiwi fee action")

	return &activityx.Action{
		Type:     typex.TransactionTransfer,
		Platform: w.Platform(),
		From:     from.String(),
		To:       to.String(),
		Metadata: metadata.TransactionTransfer(*tokenMetadata),
	}, nil
}

// NewWorker creates a new KiwiStand worker.
func NewWorker(config *config.Module) (engine.Worker, error) {
	zap.L().Debug("initializing kiwistand worker",
		zap.String("ID", config.ID),
		zap.String("network", config.Network.String()),
		zap.String("worker", config.Worker.Name()),
		zap.String("endpoint", config.Endpoint.URL),
		zap.Any("params", config.Parameters))

	var (
		err      error
		instance = worker{
			config: config,
		}
	)

	// Initialize ethereum client.
	if instance.ethereumClient, err = ethereum.Dial(context.Background(), config.Endpoint.URL, config.Endpoint.BuildEthereumOptions()...); err != nil {
		return nil, fmt.Errorf("initialize ethereum client: %w", err)
	}

	// Initialize token client.
	instance.tokenClient = token.NewClient(instance.ethereumClient)

	// Initialize kiwistand filterers.
	instance.kiwiFilterer = lo.Must(kiwistand.NewKiwiFilterer(ethereum.AddressGenesis, nil))
	instance.protocolRewardsFilterer = lo.Must(kiwistand.NewProtocolRewardsFilterer(ethereum.AddressGenesis, nil))

	zap.L().Info("kiwistand worker initialized successfully")

	return &instance, nil
}
