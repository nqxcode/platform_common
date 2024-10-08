package redis

import (
	"context"
	"log"
	"sort"
	"time"

	"github.com/nqxcode/platform_common/client/cache"

	"github.com/gomodule/redigo/redis"
)

var _ cache.RedisClient = (*client)(nil)

const (
	DefaultScanCount = 100
	ZeroCursor       = "0"
)

type handler func(ctx context.Context, conn redis.Conn) error

type client struct {
	pool   *redis.Pool
	config cache.RedisConfig
}

func NewClient(pool *redis.Pool, config cache.RedisConfig) *client {
	return &client{
		pool:   pool,
		config: config,
	}
}

// HashSet set values for key
func (c *client) HashSet(ctx context.Context, key string, values interface{}) error {
	err := c.execute(ctx, func(ctx context.Context, conn redis.Conn) error {
		_, err := conn.Do("HSET", redis.Args{key}.AddFlat(values)...)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// Set value by key
func (c *client) Set(ctx context.Context, key string, value interface{}) error {
	err := c.execute(ctx, func(ctx context.Context, conn redis.Conn) error {
		_, err := conn.Do("SET", redis.Args{key}.Add(value)...)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// HGetAll get fields by key
func (c *client) HGetAll(ctx context.Context, key string) ([]interface{}, error) {
	var values []interface{}
	err := c.execute(ctx, func(ctx context.Context, conn redis.Conn) error {
		var errEx error
		values, errEx = redis.Values(conn.Do("HGETALL", key))
		if errEx != nil {
			return errEx
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}

// MultiHGetAll multi get fields by keys
func (c *client) MultiHGetAll(ctx context.Context, keys []string) ([]cache.Values, error) {
	valuesList := make([]cache.Values, 0, len(keys))
	err := c.execute(ctx, func(ctx context.Context, conn redis.Conn) error {
		for _, key := range keys {
			err := conn.Send("HGETALL", key)
			if err != nil {
				return err
			}
		}

		err := conn.Flush()
		if err != nil {
			return err
		}

		for _, key := range keys {
			values, errEx := redis.Values(conn.Receive())
			if errEx != nil {
				return errEx
			}
			valuesList = append(valuesList, cache.Values{Key: key, Values: values})
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return valuesList, nil
}

// Get by key
func (c *client) Get(ctx context.Context, key string) (interface{}, error) {
	var value interface{}
	err := c.execute(ctx, func(ctx context.Context, conn redis.Conn) error {
		var errEx error
		value, errEx = conn.Do("GET", key)
		if errEx != nil {
			return errEx
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return value, nil
}

// Delete by key
func (c *client) Delete(ctx context.Context, key string) error {
	err := c.execute(ctx, func(ctx context.Context, conn redis.Conn) error {
		var errEx error
		_, errEx = conn.Do("DEL", key)
		if errEx != nil {
			return errEx
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// Expire value for key
func (c *client) Expire(ctx context.Context, key string, expiration time.Duration) error {
	err := c.execute(ctx, func(ctx context.Context, conn redis.Conn) error {
		_, err := conn.Do("EXPIRE", key, int(expiration.Seconds()))
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// Ping ping redis
func (c *client) Ping(ctx context.Context) error {
	err := c.execute(ctx, func(ctx context.Context, conn redis.Conn) error {
		_, err := conn.Do("PING")
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// FlushDB flush redis db
func (c *client) FlushDB(ctx context.Context) error {
	err := c.execute(ctx, func(ctx context.Context, conn redis.Conn) error {
		_, err := conn.Do("FLUSHDB")
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// Scan keys with pattern
func (c *client) Scan(ctx context.Context, pattern string, options ...cache.ScanOption) ([]string, error) {
	opts := cache.ScanOptions{
		ScanCount: DefaultScanCount,
	}
	for _, o := range options {
		o(&opts)
	}

	result := make([]string, 0)

	cursor := ZeroCursor
	count := opts.ScanCount

	err := c.execute(ctx, func(ctx context.Context, conn redis.Conn) error {
		for {
			values, err := redis.Values(conn.Do("SCAN", cursor, "MATCH", pattern, "COUNT", count))
			if err != nil {
				return err
			}

			cursor, _ = redis.String(values[0], nil)
			keys, _ := redis.Strings(values[1], nil)

			for _, key := range keys {
				result = append(result, key)
			}

			if cursor == ZeroCursor {
				break
			}
		}

		if opts.KeyComparator != nil {
			sort.Slice(result, func(i, j int) bool {
				return (*opts.KeyComparator)(result[i], result[j])
			})
		}

		return nil
	})

	return result, err
}

// RPush push value to end of list by key
func (c *client) RPush(ctx context.Context, key string, values []interface{}) error {
	err := c.execute(ctx, func(ctx context.Context, conn redis.Conn) error {
		_, err := conn.Do("RPUSH", redis.Args{key}.AddFlat(values)...)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

// LRange range of values from list by key
func (c *client) LRange(ctx context.Context, key string, start, stop int) ([]interface{}, error) {
	var values []interface{}
	err := c.execute(ctx, func(ctx context.Context, conn redis.Conn) error {
		var errEx error
		values, errEx = redis.Values(conn.Do("LRANGE", redis.Args{key}.Add(start).Add(stop)...))
		if errEx != nil {
			return errEx
		}

		return nil
	})

	return values, err
}

func (c *client) execute(ctx context.Context, handler handler) error {
	conn, err := c.getConnect(ctx)
	if err != nil {
		return err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Printf("failed to close redis connection: %v\n", err)
		}
	}()

	err = handler(ctx, conn)
	if err != nil {
		return err
	}

	return nil
}

func (c *client) getConnect(ctx context.Context) (redis.Conn, error) {
	getConnTimeoutCtx, cancel := context.WithTimeout(ctx, c.config.ConnectionTimeout())
	defer cancel()

	conn, err := c.pool.GetContext(getConnTimeoutCtx)
	if err != nil {
		log.Printf("failed to get redis connection: %v\n", err)

		_ = conn.Close()
		return nil, err
	}

	return conn, nil
}
