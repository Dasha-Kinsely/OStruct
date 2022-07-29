package config

import (
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v9"
)

func SetupRedisConnection() *redis.Client {
	// does the current version require caching features
	useCache, parseErr := strconv.ParseBool(os.Getenv("USE_CACHE"))
	if parseErr != nil {
		panic("Error parsing boolean variables from .env files, please check your .env files...")
	}
	if !useCache {
		return nil
	}
	// initialize redis for caching
	addr := os.Getenv("REDIS_ADDR")
	password := os.Getenv("REDIS_PASSWORD")
	rdb, parseErr := strconv.Atoi(os.Getenv("REDIS_DB"))
	poolSize, parseErr := strconv.Atoi(os.Getenv("REDIS_POOL_SIZE"))
	minIdleConns, parseErr := strconv.Atoi(os.Getenv("REDIS_MIN_IDLE_CONN"))
	if parseErr != nil {
		panic("Error parsing boolean variables from .env files, please check your .env files...")
	}
	RedisClient := redis.NewClient(&redis.Options{
		Addr: addr,
		Password: password,
		DB: rdb,
		PoolSize: poolSize,
		MaxRetries: 3,
		DialTimeout: 5*time.Second,
		ReadTimeout: 5*time.Second,
		WriteTimeout: 5*time.Second,
		MinIdleConns: minIdleConns,
	})
	return RedisClient
}
