package cache

import "time"

// RedisConfig redis config
type RedisConfig interface {
	Address() string
	Password() string
	ConnectionTimeout() time.Duration
	MaxIdle() int
	IdleTimeout() time.Duration
	DB() int
}
