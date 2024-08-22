package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisClient struct {
	*redis.Client
}

// NewRedisClient new Redis client
func NewRedisClient(ctx context.Context, cfg *Config) (Cache, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("redis.Ping: %s", err.Error())
	}

	return &redisClient{rdb}, nil
}

// Get value from cache
func (r *redisClient) Get(ctx context.Context, key string) ([]byte, error) {
	value, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("can not get cache: %w", err)
	}

	return []byte(value), nil
}

// Set value to cache for key
func (r *redisClient) Set(ctx context.Context, key string, value any, TTLInSec *int64) error {
	err := r.Client.Set(
		ctx,
		key,
		value,
		func() time.Duration {
			if TTLInSec != nil {
				return time.Duration(*TTLInSec) * time.Second
			} else {
				return redis.KeepTTL
			}
		}(),
	).Err()
	if err != nil {
		return fmt.Errorf("can not set cache: %w", err)
	}
	return nil
}

// SetWithTags set value with tags to cache
func (r *redisClient) SetWithTags(ctx context.Context, key string, value any, tags []string, TTLInSec *int64) error {
	err := r.Client.Set(
		ctx,
		key,
		value,
		func() time.Duration {
			if TTLInSec != nil {
				return time.Duration(*TTLInSec) * time.Second
			} else {
				return redis.KeepTTL
			}
		}(),
	).Err()
	if err != nil {
		return fmt.Errorf("can not set cache: %w", err)
	}

	for _, tag := range tags {
		err = r.Client.SAdd(ctx, tag, key).Err()
		if err != nil {
			return fmt.Errorf("can not set cache: %w", err)
		}
	}

	return nil
}

// Delete delete from cache by keys
func (r *redisClient) Delete(ctx context.Context, keys ...string) error {
	if err := r.Client.Del(ctx, keys...).Err(); err != nil {
		return fmt.Errorf("can not delete cache: %w", err)
	}
	return nil
}

// DeleteByTags delete from cache by tags
func (r *redisClient) DeleteByTags(ctx context.Context, tags []string) error {
	keySet := make(map[string]struct{})

	for _, tag := range tags {
		keys, err := r.Client.SMembers(ctx, tag).Result()
		if err != nil {
			return fmt.Errorf("can not delete cache: %w", err)
		}
		for _, key := range keys {
			keySet[key] = struct{}{}
		}
	}

	for key := range keySet {
		err := r.Client.Del(ctx, key).Err()
		if err != nil {
			return fmt.Errorf("can not delete cache: %w", err)
		}
	}

	for _, tag := range tags {
		err := r.Client.Del(ctx, tag).Err()
		if err != nil {
			return fmt.Errorf("can not delete cache: %w", err)
		}
	}

	return nil
}

// Ping ping redis
func (r *redisClient) Ping(ctx context.Context) error {
	_, err := r.Client.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

// FlushDB flush redis db
func (r *redisClient) FlushDB(ctx context.Context) error {
	return r.Client.FlushDB(ctx).Err()
}
