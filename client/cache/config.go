package cache

import "time"

// RedisConfig redis config
type RedisConfig interface {
	Address() string
	ConnectionTimeout() time.Duration
	MaxIdle() int
	IdleTimeout() time.Duration
	Password() string
	DB() int
}
