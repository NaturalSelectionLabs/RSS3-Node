package cockroachdb_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/cockroachdb"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/database/dialer"
	"github.com/rss3-network/node/internal/database/model"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name            string
		driver          database.Driver
		partition       *bool
		activityCreated []*activityx.Activity
		activityUpdated []*activityx.Activity
	}{
		{
			name:      "cockroach",
			driver:    database.DriverCockroachDB,
			partition: lo.ToPtr(true),
			activityCreated: []*activityx.Activity{
				{
					ID:      "0x30182d4468ddc7001b897908203abb57939fc57663c491435a2f88cafd51d101",
					Network: network.Ethereum,
					From:    "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
					To:      "0x9D22816f6611cFcB0cDE5076C5f4e4A269E79Bef",
					Type:    typex.TransactionTransfer,
					Actions: []*activityx.Action{
						{
							Type: typex.TransactionTransfer,
							From: "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
							To:   "0x9D22816f6611cFcB0cDE5076C5f4e4A269E79Bef",
						},
					},
					Timestamp: uint64(time.Now().Unix()),
				},
				{
					ID:      "0xedc029f7c7acce7b72939f8bfff44c5fdc7e64e3e2ba650d195799db8fec4c90",
					Network: network.Ethereum,
					From:    "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
					To:      "0x9D22816f6611cFcB0cDE5076C5f4e4A269E79Bef",
					Type:    typex.TransactionTransfer,
					Actions: []*activityx.Action{
						{
							Type: typex.TransactionTransfer,
							From: "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
							To:   "0x9D22816f6611cFcB0cDE5076C5f4e4A269E79Bef",
						},
					},
					Timestamp: uint64(time.Now().Add(-3 * 31 * 24 * time.Hour).Unix()),
				},
			},
			activityUpdated: []*activityx.Activity{
				{
					ID:      "0x30182d4468ddc7001b897908203abb57939fc57663c491435a2f88cafd51d101",
					Network: network.Ethereum,
					From:    "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
					To:      "0x9D22816f6611cFcB0cDE5076C5f4e4A269E79Bef",
					Type:    typex.TransactionTransfer,
					Actions: []*activityx.Action{
						{
							Type: typex.TransactionTransfer,
							From: "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
							To:   "0x9D22816f6611cFcB0cDE5076C5f4e4A269E79Bef",
						},
						{
							Type: typex.TransactionTransfer,
							From: "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
							To:   "0x9D22816f6611cFcB0cDE5076C5f4e4A269E79Bef",
						},
					},
					Timestamp: uint64(time.Now().Unix()),
				},
			},
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			var (
				container      *gnomock.Container
				dataSourceName string
				err            error
			)

			for {
				container, dataSourceName, err = createContainer(context.Background(), testcase.driver, *testcase.partition)
				if err == nil {
					break
				}
			}

			t.Cleanup(func() {
				require.NoError(t, gnomock.Stop(container))
			})

			// Dial the database.
			client, err := dialer.Dial(context.Background(), &config.Database{
				Driver:    testcase.driver,
				URI:       dataSourceName,
				Partition: testcase.partition,
			})

			require.NoError(t, err)
			require.NotNil(t, client)

			// Migrate the database.
			require.NoError(t, client.Migrate(context.Background()))

			// Begin a transaction.
			require.NoError(t, client.SaveActivities(context.Background(), testcase.activityCreated))

			// Query first activity
			for _, activity := range testcase.activityCreated {
				data, page, err := client.FindActivity(context.Background(), model.ActivityQuery{ID: lo.ToPtr(activity.ID), ActionLimit: 10})
				require.NoError(t, err)
				require.NotNil(t, data)
				require.Greater(t, lo.FromPtr(page), 0)
				require.Equal(t, data.ID, activity.ID)
				require.Equal(t, data.From, activity.From)
				require.Equal(t, data.To, activity.To)
			}

			// Query activities by account.
			accounts := make(map[string]int)

			for _, activity := range testcase.activityCreated {
				accounts[activity.From]++
				accounts[activity.To]++
			}

			for account, count := range accounts {
				activities, err := client.FindActivities(context.Background(), model.ActivitiesQuery{Owner: &account, Limit: 100})
				require.NoError(t, err)
				require.Len(t, activities, count)
			}

			// Update activities.
			require.NoError(t, client.SaveActivities(context.Background(), testcase.activityUpdated))
		})
	}
}

func createContainer(ctx context.Context, driver database.Driver, partition bool) (container *gnomock.Container, dataSourceName string, err error) {
	config := config.Database{
		Driver:    driver,
		Partition: &partition,
	}

	switch driver {
	case database.DriverCockroachDB:
		preset := cockroachdb.Preset(
			cockroachdb.WithDatabase("test"),
			cockroachdb.WithVersion("v23.1.8"),
		)

		// Use a health check function to wait for the database to be ready.
		healthcheckFunc := func(ctx context.Context, container *gnomock.Container) error {
			config.URI = formatContainerURI(container)

			client, err := dialer.Dial(ctx, &config)
			if err != nil {
				return err
			}

			transaction, err := client.Begin(ctx)
			if err != nil {
				return err
			}

			defer lo.Try(transaction.Rollback)

			return nil
		}

		container, err = gnomock.Start(preset, gnomock.WithContext(ctx), gnomock.WithHealthCheck(healthcheckFunc))
		if err != nil {
			return nil, "", err
		}

		return container, formatContainerURI(container), nil
	default:
		return nil, "", fmt.Errorf("unsupported driver: %s", driver)
	}
}

func formatContainerURI(container *gnomock.Container) string {
	return fmt.Sprintf(
		"postgres://root@%s:%d/%s?sslmode=disable",
		container.Host,
		container.DefaultPort(),
		"test",
	)
}
