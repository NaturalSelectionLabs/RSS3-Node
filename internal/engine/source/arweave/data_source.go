package arweave

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/provider/arweave"
	"github.com/rss3-network/node/provider/arweave/bundle"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
	"github.com/sourcegraph/conc/pool"
	"go.uber.org/zap"
)

const (
	// The block time in Arweave mainnet is designed to be approximately 2 minutes.
	defaultBlockTime = 120 * time.Second
)

// Ensure that dataSource implements DataSource.
var _ engine.DataSource = (*dataSource)(nil)

type dataSource struct {
	config        *config.Module
	option        *Option
	filter        *Filter
	arweaveClient arweave.Client
	state         State
}

func (s *dataSource) Network() network.Network {
	return s.config.Network
}

func (s *dataSource) State() json.RawMessage {
	return lo.Must(json.Marshal(s.state))
}

func (s *dataSource) Start(ctx context.Context, tasksChan chan<- *engine.Tasks, errorChan chan<- error) {
	// Initialize dataSource.
	if err := s.initialize(); err != nil {
		errorChan <- fmt.Errorf("initialize dataSource: %w", err)

		return
	}

	// Start a goroutine to poll blocks.
	go func() {
		retryableFunc := func() error {
			if err := s.pollBlocks(ctx, tasksChan, s.filter); err != nil {
				return fmt.Errorf("poll blocks: %w", err)
			}

			return nil
		}

		err := retry.Do(retryableFunc,
			retry.Attempts(0),
			retry.Delay(time.Second),
			retry.DelayType(retry.BackOffDelay),
			retry.MaxDelay(5*time.Minute),
			retry.OnRetry(func(n uint, err error) {
				zap.L().Error("retry arweave dataSource start", zap.Uint("retry", n), zap.Error(err))
			}),
		)
		if err != nil {
			errorChan <- err
		}
	}()
}

// initialize initializes the dataSource.
func (s *dataSource) initialize() (err error) {
	// Initialize arweave client.
	if s.arweaveClient, err = arweave.NewClient(); err != nil {
		return fmt.Errorf("create arweave client: %w", err)
	}

	return nil
}

// initializeBlockHeights initializes block heights.
func (s *dataSource) initializeBlockHeights() {
	if s.option.BlockStart != nil && s.option.BlockStart.Uint64() > s.state.BlockHeight {
		s.state.BlockHeight = s.option.BlockStart.Uint64()
	}
}

// pollBlocks polls blocks from arweave network.
func (s *dataSource) pollBlocks(ctx context.Context, tasksChan chan<- *engine.Tasks, filter *Filter) error {
	var (
		blockHeightLatestRemote int64
		err                     error
	)

	// Get start block height from config
	// if not set, use default value 0
	s.initializeBlockHeights()

	// Get target block height from config
	// if not set, use the latest block height from arweave network
	if s.option.BlockTarget != nil {
		zap.L().Info("block height target", zap.Uint64("block.height.target", s.option.BlockTarget.Uint64()))
		blockHeightLatestRemote = int64(s.option.BlockTarget.Uint64())
	} else {
		// Get remote block height from arweave network.
		blockHeightLatestRemote, err = s.arweaveClient.GetBlockHeight(ctx)
		if err != nil {
			return fmt.Errorf("get latest block height: %w", err)
		}

		zap.L().Info("get latest block height", zap.Int64("block.height", blockHeightLatestRemote))
	}

	for {
		if s.option.BlockTarget != nil && s.option.BlockTarget.Uint64() <= s.state.BlockHeight {
			break
		}

		// Check if block height is latest.
		if s.state.BlockHeight >= uint64(blockHeightLatestRemote) {
			// Get the latest block height from arweave network for reconfirming.
			if blockHeightLatestRemote, err = s.arweaveClient.GetBlockHeight(ctx); err != nil {
				return fmt.Errorf("get latest block height: %w", err)
			}

			zap.L().Info("get latest block height", zap.Int64("block.height", blockHeightLatestRemote))

			if s.state.BlockHeight >= uint64(blockHeightLatestRemote) {
				// Wait for the next block on arweave network.
				time.Sleep(defaultBlockTime)
			}

			continue
		}

		// If the block height of state is 0, force to start from 0.
		blockHeightStart := lo.Ternary(s.state.BlockHeight == 0, 0, s.state.BlockHeight+1)

		// Pull blocks
		blockHeightEnd := lo.Min([]uint64{
			uint64(blockHeightLatestRemote),
			blockHeightStart + *s.option.ConcurrentBlockRequests,
		})

		// Pull blocks by range.
		blocks, err := s.batchPullBlocksByRange(ctx, blockHeightStart, blockHeightEnd)
		if err != nil {
			return fmt.Errorf("batch pull blocks: %w", err)
		}

		// Pull transactions.
		transactionIDs := lo.FlatMap(blocks, func(block *arweave.Block, _ int) []string {
			return block.Txs
		})

		// Batch pull transactions by ids.
		transactions, err := s.batchPullTransactions(ctx, transactionIDs)
		if err != nil {
			return fmt.Errorf("batch pull transactions: %w", err)
		}

		// Filter transactions by owner.
		transactions = s.filterOwnerTransaction(transactions, append(filter.OwnerAddresses, filter.BundlrAddresses...))

		// Pull transaction data and ignore the transaction if the owner is bundlr node.
		if err := s.batchPullData(ctx, transactions); err != nil {
			return fmt.Errorf("batch pull data: %w", err)
		}

		// Decode Bundle transactions group by block.
		for index, block := range blocks {
			bundleTransactionIDs := s.GroupBundleTransactions(transactions, block)

			bundleTransactions, err := s.batchPullBundleTransactions(ctx, bundleTransactionIDs)
			if err != nil {
				return fmt.Errorf("pull bundle transacctions: %w", err)
			}

			for _, bundleTransaction := range bundleTransactions {
				blocks[index].Txs = append(blocks[index].Txs, bundleTransaction.ID)
			}

			transactions = append(transactions, bundleTransactions...)
		}

		// Discard the Bundle transaction itself.
		transactions = s.discardRootBundleTransaction(transactions)

		// Discard duplicate bundle transactions.
		// https://viewblock.io/arweave/block/1187748 has duplicate bundle transactions.
		//
		// $ sha1sum 4mdtwXkR3V9qzA2haO0TG2mgl2bhanROywKPVu6QkCQ fnsyKm1hw4xSFqXkmJ4HzPrK8wZnlpEJjGcDDn3iXvI
		// 225c6bcb20b39b1557c80fa88ff3960dcc901031  4mdtwXkR3V9qzA2haO0TG2mgl2bhanROywKPVu6QkCQ
		// 225c6bcb20b39b1557c80fa88ff3960dcc901031  fnsyKm1hw4xSFqXkmJ4HzPrK8wZnlpEJjGcDDn3iXvI
		transactions = s.discardDuplicateBundleTransaction(transactions)

		tasks := s.buildTasks(ctx, blocks, transactions)

		// TODO It might be possible to use generics to avoid manual type assertions.
		tasksChan <- tasks

		// Update block height to state.
		s.state.BlockHeight = blockHeightEnd
	}

	return nil
}

// batchPullBlocksByRange pulls blocks by range, from local state block height to remote block height.
func (s *dataSource) batchPullBlocksByRange(ctx context.Context, blockHeightStart, blockHeightEnd uint64) ([]*arweave.Block, error) {
	zap.L().Info("begin to batch pull transactions by range", zap.Uint64("block.height.start", blockHeightStart), zap.Uint64("block.height.end", blockHeightEnd))

	// Pull blocks by block heights.
	blockHeights := lo.Map(lo.RangeWithSteps(blockHeightStart, blockHeightEnd+1, 1), func(blockHeight uint64, _ int) *big.Int {
		return new(big.Int).SetUint64(blockHeight)
	})

	blocks, err := s.batchPullBlocks(ctx, blockHeights)
	if err != nil {
		return nil, fmt.Errorf("batch pull blocks: %w", err)
	}

	return blocks, nil
}

// batchPullBlocks pulls blocks by block heights.
func (s *dataSource) batchPullBlocks(ctx context.Context, blockHeights []*big.Int) ([]*arweave.Block, error) {
	zap.L().Info("begin to pull blocks", zap.Int("blocks", len(blockHeights)))

	resultPool := pool.NewWithResults[*arweave.Block]().
		WithContext(ctx).
		WithCancelOnError()

	for _, blockHeight := range blockHeights {
		blockHeight := blockHeight

		resultPool.Go(func(ctx context.Context) (*arweave.Block, error) {
			retryableFunc := func() (*arweave.Block, error) {
				return s.arweaveClient.GetBlockByHeight(ctx, blockHeight.Int64())
			}

			return retry.DoWithData(
				retryableFunc,
				retry.Attempts(defaultRetryAttempts),
				retry.Delay(defaultRetryDelay),
				retry.DelayType(retry.BackOffDelay),
				retry.OnRetry(func(attempt uint, err error) {
					zap.L().Error("retry pull block", zap.Stringer("block.height", blockHeight), zap.Uint("attempt", attempt), zap.Error(err))
				}),
			)
		})
	}

	return resultPool.Wait()
}

// batchPullTransactions pulls transactions by transaction ids.
func (s *dataSource) batchPullTransactions(ctx context.Context, transactionIDs []string) ([]*arweave.Transaction, error) {
	zap.L().Info("begin to pull transactions", zap.Int("transactions", len(transactionIDs)))

	resultPool := pool.NewWithResults[*arweave.Transaction]().
		WithContext(ctx).
		WithCancelOnError().
		WithMaxGoroutines(int(lo.FromPtr(s.option.ConcurrentBlockRequests)))

	for _, transactionID := range transactionIDs {
		transactionID := transactionID

		resultPool.Go(func(ctx context.Context) (*arweave.Transaction, error) {
			retryableFunc := func() (*arweave.Transaction, error) {
				return s.arweaveClient.GetTransactionByID(ctx, transactionID)
			}

			return retry.DoWithData(
				retryableFunc,
				retry.Attempts(defaultRetryAttempts),
				retry.Delay(defaultRetryDelay),
				retry.DelayType(retry.BackOffDelay),
				retry.OnRetry(func(attempt uint, err error) {
					zap.L().Error("retry pull transaction", zap.String("transaction.id", transactionID), zap.Uint("attempt", attempt), zap.Error(err))
				}),
			)
		})
	}

	return resultPool.Wait()
}

// batchPullData pulls data by transactions.
// It will discard the transaction if the owner is bundlr node.
func (s *dataSource) batchPullData(ctx context.Context, transactions []*arweave.Transaction) error {
	resultPool := pool.New().
		WithContext(ctx).
		WithCancelOnError().
		WithMaxGoroutines(int(lo.FromPtr(s.option.ConcurrentBlockRequests)))

	for index, transaction := range transactions {
		index, transaction := index, transaction

		owner, err := arweave.PublicKeyToAddress(transaction.Owner)
		if err != nil {
			return fmt.Errorf("invalid owner of transaction %s: %w", transaction.ID, err)
		}

		if lo.Contains(s.filter.BundlrAddresses, owner) {
			continue
		}

		resultPool.Go(func(ctx context.Context) error {
			retryableFunc := func() (string, error) {
				response, err := s.arweaveClient.GetTransactionData(ctx, transaction.ID)
				if err != nil {
					if errors.Is(err, arweave.ErrorNotFound) {
						return "", nil
					}

					return "", fmt.Errorf("fetch transaction %s data: %w", transaction.ID, err)
				}

				defer lo.Try(response.Close)

				buffer := new(bytes.Buffer)
				if _, err := io.Copy(base64.NewEncoder(base64.RawURLEncoding, buffer), response); err != nil {
					return "", fmt.Errorf("read and encode response: %w", err)
				}

				return buffer.String(), nil
			}

			transactions[index].Data, err = retry.DoWithData(
				retryableFunc,
				retry.Attempts(defaultRetryAttempts),
				retry.Delay(defaultRetryDelay),
				retry.DelayType(retry.BackOffDelay),
				retry.OnRetry(func(attempt uint, err error) {
					zap.L().Error("retry pull transaction data", zap.String("transaction.id", transaction.ID), zap.Uint("attempt", attempt), zap.Error(err))
				}),
			)

			return err
		})
	}

	return resultPool.Wait()
}

// batchPullBundleTransactions pulls bundle transactions by transaction ids.
func (s *dataSource) batchPullBundleTransactions(ctx context.Context, transactionIDs []string) ([]*arweave.Transaction, error) {
	zap.L().Info("begin to pull and filter bundle transactions", zap.Int("transactions", len(transactionIDs)))

	resultPool := pool.NewWithResults[[]*arweave.Transaction]().
		WithContext(ctx).
		WithFirstError().
		WithCancelOnError()

	for _, transactionID := range transactionIDs {
		transactionID := transactionID

		resultPool.Go(func(ctx context.Context) ([]*arweave.Transaction, error) {
			retryableFunc := func() ([]*arweave.Transaction, error) {
				bundleTransactions, err := s.pullBundleTransactions(ctx, transactionID)
				if err != nil {
					return nil, fmt.Errorf("get bundle transaction %s: %w", transactionID, err)
				}

				return bundleTransactions, nil
			}

			return retry.DoWithData(
				retryableFunc,
				retry.Attempts(defaultRetryAttempts),
				retry.Delay(defaultRetryDelay),
				retry.DelayType(retry.BackOffDelay),
				retry.OnRetry(func(attempt uint, err error) {
					zap.L().Error("retry pull bundle transaction", zap.String("transaction.id", transactionID), zap.Uint("attempt", attempt), zap.Error(err))
				}),
			)
		})
	}

	bundleTransactions, err := resultPool.Wait()
	if err != nil {
		return nil, fmt.Errorf("wait result pool: %w", err)
	}

	return lo.Flatten(bundleTransactions), nil
}

// pullBundleTransactions pulls bundle transactions by transaction id.
func (s *dataSource) pullBundleTransactions(ctx context.Context, transactionID string) ([]*arweave.Transaction, error) {
	bundleTransactions := make([]*arweave.Transaction, 0)

	response, err := s.arweaveClient.GetTransactionData(ctx, transactionID)
	if err != nil {
		if errors.Is(err, arweave.ErrorNotFound) {
			return nil, nil
		}

		return nil, fmt.Errorf("fetch transaction: %w", err)
	}

	defer lo.Try(response.Close)

	decoder := bundle.NewDecoder(response)

	header, err := decoder.DecodeHeader()
	if err != nil {
		// Ignore invalid bundle transaction.
		zap.L().Error("discard a invalid bundle transaction", zap.String("transaction_id", transactionID))

		return nil, nil
	}

	for index := 0; decoder.Next(); index++ {
		dataItemInfo := header.DataItemInfos[index]

		dataItem, err := decoder.DecodeDataItem()
		if err != nil {
			// Ignore invalid signature and data length.
			zap.L().Error("decode data item", zap.Error(err), zap.String("transaction_id", transactionID))

			return nil, nil
		}

		bundleTransaction := arweave.Transaction{
			Format: 2,
			ID:     dataItemInfo.ID,
			Owner:  dataItem.Owner,
			Tags: lo.Map(dataItem.Tags, func(tag bundle.Tag, _ int) arweave.Tag {
				return arweave.Tag{
					Name:  arweave.Base64Encode(tag.Name),
					Value: arweave.Base64Encode(tag.Value),
				}
			}),
			Target:    dataItem.Target,
			Signature: dataItem.Signature,
		}

		transactionOwner, err := arweave.PublicKeyToAddress(bundleTransaction.Owner)
		if err != nil {
			zap.L().Error("invalid owner of transaction", zap.String("id", dataItemInfo.ID), zap.Any("owner", bundleTransaction.Owner), zap.Error(err))

			continue
		}

		// Filter owner addresses.
		if !lo.Contains(s.filter.OwnerAddresses, transactionOwner) {
			if _, err := io.Copy(io.Discard, dataItem); err != nil {
				return nil, fmt.Errorf("discard data item %s: %w", dataItemInfo.ID, err)
			}

			continue
		}

		data, err := io.ReadAll(dataItem)
		if err != nil {
			return nil, fmt.Errorf("read data item %s: %w", dataItemInfo.ID, err)
		}

		bundleTransaction.Data = arweave.Base64Encode(data)
		bundleTransaction.DataSize = strconv.Itoa(len(bundleTransaction.Data))

		bundleTransactions = append(bundleTransactions, &bundleTransaction)
	}

	return bundleTransactions, nil
}

// GroupBundleTransactions groups bundle transactions by block.
func (s *dataSource) GroupBundleTransactions(transactions []*arweave.Transaction, block *arweave.Block) []string {
	return lo.FilterMap(transactions, func(transaction *arweave.Transaction, _ int) (string, bool) {
		hasBundleFormatTag := lo.ContainsBy(transaction.Tags, func(tag arweave.Tag) bool {
			tagName, err := base64.RawURLEncoding.DecodeString(tag.Name)
			if err != nil {
				return false
			}

			tagValue, err := base64.RawURLEncoding.DecodeString(tag.Value)
			if err != nil {
				return false
			}

			return strings.EqualFold(string(tagName), "Bundle-Format") && strings.EqualFold(string(tagValue), "binary")
		})

		hasBundleVersionTag := lo.ContainsBy(transaction.Tags, func(tag arweave.Tag) bool {
			tagName, err := base64.RawURLEncoding.DecodeString(tag.Name)
			if err != nil {
				return false
			}

			tagValue, err := base64.RawURLEncoding.DecodeString(tag.Value)
			if err != nil {
				return false
			}

			return strings.EqualFold(string(tagName), "Bundle-Version") && strings.EqualFold(string(tagValue), "2.0.0")
		})

		if !(hasBundleFormatTag && hasBundleVersionTag) {
			return "", false
		}

		if !lo.Contains(block.Txs, transaction.ID) {
			return "", false
		}

		owner, err := arweave.PublicKeyToAddress(transaction.Owner)
		if err != nil {
			zap.L().Error("invalid owner of transaction", zap.String("transaction_id", transaction.ID), zap.Error(err))

			return "", false
		}

		return transaction.ID, lo.Contains(s.filter.BundlrAddresses, owner)
	})
}

// discardRootBundleTransaction discards the root bundle transaction.
func (s *dataSource) discardRootBundleTransaction(transactions []*arweave.Transaction) []*arweave.Transaction {
	return lo.Filter(transactions, func(transaction *arweave.Transaction, _ int) bool {
		transactionOwner, err := arweave.PublicKeyToAddress(transaction.Owner)
		if err != nil {
			return false
		}

		return !lo.Contains(s.filter.BundlrAddresses, transactionOwner)
	})
}

// discardDuplicateBundleTransaction discards duplicate bundle transactions.
func (s *dataSource) discardDuplicateBundleTransaction(transactions []*arweave.Transaction) []*arweave.Transaction {
	var (
		cache   = make(map[string]struct{})
		results = make([]*arweave.Transaction, 0, len(transactions))
	)

	for index := range transactions {
		if _, found := cache[transactions[index].ID]; found {
			continue
		}

		cache[transactions[index].ID] = struct{}{}

		results = append(results, transactions[index])
	}

	return results
}

// filterOwnerTransaction filters owner transactions.
func (s *dataSource) filterOwnerTransaction(transactions []*arweave.Transaction, ownerAddress []string) []*arweave.Transaction {
	return lo.Filter(transactions, func(transaction *arweave.Transaction, _ int) bool {
		transactionOwner, err := arweave.PublicKeyToAddress(transaction.Owner)
		if err != nil {
			return false
		}

		return lo.Contains(ownerAddress, transactionOwner)
	})
}

// buildTasks builds tasks from blocks and transactions.
func (s *dataSource) buildTasks(_ context.Context, blocks []*arweave.Block, transactions []*arweave.Transaction) *engine.Tasks {
	var tasks engine.Tasks

	for _, transaction := range transactions {
		block, _ := lo.Find(blocks, func(block *arweave.Block) bool {
			return lo.Contains(block.Txs, transaction.ID)
		})

		tasks.Tasks = append(tasks.Tasks, &Task{
			Network:     s.Network(),
			Block:       *block,
			Transaction: *transaction,
		})
	}

	return &tasks
}

// NewSource creates a new arweave dataSource.
func NewSource(config *config.Module, sourceFilter engine.DataSourceFilter, checkpoint *engine.Checkpoint) (engine.DataSource, error) {
	var (
		state State
		err   error
	)

	// Initialize state from checkpoint.
	if checkpoint != nil {
		if err := json.Unmarshal(checkpoint.State, &state); err != nil {
			return nil, err
		}
	}

	instance := dataSource{
		config: config,
		filter: new(Filter), // Set a default filter for the dataSource.
		state:  state,
	}

	// Initialize filter.
	if sourceFilter != nil {
		var ok bool
		if instance.filter, ok = sourceFilter.(*Filter); !ok {
			return nil, fmt.Errorf("invalid dataSource filter type %T", sourceFilter)
		}
	}

	if instance.option, err = NewOption(config.Network, config.Parameters); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return &instance, nil
}