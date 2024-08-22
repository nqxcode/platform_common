package cache

import (
	"context"
)

type Cache interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value any, TTLInSec *int64) error
	SetWithTags(ctx context.Context, key string, value any, tags []string, TTLInSec *int64) error
	Delete(ctx context.Context, keys ...string) error
	DeleteByTags(ctx context.Context, tags []string) error
	Ping(ctx context.Context) error
	FlushDB(ctx context.Context) error
}
