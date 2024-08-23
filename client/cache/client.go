package cache

import (
	"context"
	"time"
)

// RedisClient redis client interface
type RedisClient interface {
	HashSet(ctx context.Context, key string, values interface{}) error
	Set(ctx context.Context, key string, value interface{}) error
	HGetAll(ctx context.Context, key string) ([]interface{}, error)
	Get(ctx context.Context, key string) (interface{}, error)
	Delete(ctx context.Context, key string) error
	Expire(ctx context.Context, key string, expiration time.Duration) error
	Ping(ctx context.Context) error
	FlushDB(ctx context.Context) error
	Scan(ctx context.Context, pattern string, options ...ScanOption) ([]string, error)
	RPush(ctx context.Context, key string, values []interface{}) error
	LRange(ctx context.Context, key string, start, stop int) ([]interface{}, error)
}

// KeyComparator comparator function type
type KeyComparator func(a, b string) bool

// ScanOptions scan options
type ScanOptions struct {
	KeyComparator *KeyComparator
}

// ScanOption scan option
type ScanOption func(*ScanOptions)

// WithKeyComparator key comparator function
func WithKeyComparator(keyComparator KeyComparator) ScanOption {
	return func(options *ScanOptions) {
		options.KeyComparator = &keyComparator
	}
}
