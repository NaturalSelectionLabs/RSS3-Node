package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/rueidis"
)

const handleUpdatesKey = "handle_updates"

// GetAllHandles retrieves all handles from the sorted set without filtering by score
func GetAllHandles(ctx context.Context, client rueidis.Client) ([]string, error) {
	cmd := client.B().Zrange().Key(handleUpdatesKey).Min("-inf").Max("+inf").Build()
	result, err := client.Do(ctx, cmd).AsStrSlice()

	if err != nil {
		return nil, fmt.Errorf("failed to get all handles: %w", err)
	}

	return result, nil
}

// AddHandleUpdate adds a handle to the sorted set with the current timestamp as score
func AddHandleUpdate(ctx context.Context, client rueidis.Client, handle string) error {
	fmt.Println("Reached AddHandleUpdate!")

	err := client.Do(ctx, client.B().Set().Key("test_key").Value("test_value").Build()).Error()
	if err != nil {
		fmt.Println("Failed to connect to Redis")
	} else {
		fmt.Println("Connected to Redis successfully")
	}

	cmd := client.B().Zadd().Key(handleUpdatesKey).ScoreMember().ScoreMember(float64(time.Now().Unix()), handle).Build()

	return client.Do(ctx, cmd).Error()
}

// GetRecentHandleUpdates retrieves handles updated since the given time
func GetRecentHandleUpdates(ctx context.Context, client rueidis.Client, since time.Time) ([]string, error) {
	cmd := client.B().Zrangebyscore().Key(handleUpdatesKey).Min(fmt.Sprintf("%d", since.Unix())).Max("+inf").Build()
	result, err := client.Do(ctx, cmd).AsStrSlice()

	if err != nil {
		return nil, err
	}

	return result, nil
}

// RemoveOldHandleUpdates removes handles updated before the given time
func RemoveOldHandleUpdates(ctx context.Context, client rueidis.Client, before time.Time) error {
	cmd := client.B().Zremrangebyscore().Key(handleUpdatesKey).Min("-inf").Max(fmt.Sprintf("%d", before.Unix())).Build()
	return client.Do(ctx, cmd).Error()
}
