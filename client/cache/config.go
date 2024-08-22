package cache

// RedisConfig redis config
type RedisConfig interface {
	Address() string
	Password() string
	DB() int
}
