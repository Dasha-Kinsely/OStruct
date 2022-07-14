package config

import (
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v9"
)

var RedisClient *redis.Client

func SetupRedisConnection() {
	// does the current version require caching features
	useCache, parseErr := strconv.ParseBool(os.Getenv("USE_CACHE"))
	if parseErr != nil {
		panic("Error parsing boolean variables from .env files, please check your .env files...")
	}
	if !useCache {
		return
	}
	// initialize redis for caching
	addr := os.Getenv("REDIS_ADDR")
	password := os.Getenv("REDIS_PASSWORD")
	poolSize, parseErr := strconv.Atoi(os.Getenv("REDIS_POOLSIZE"))
	minIdleConns, parseErr := strconv.Atoi(os.Getenv("REDIS_MINIDLECONNS"))
	if parseErr != nil {
		panic("Error parsing boolean variables from .env files, please check your .env files...")
	}
	RedisClient = redis.NewClient(&redis.Options{
		Addr: addr,
		Password: password,
		DB: 0,
		PoolSize: poolSize,
		MaxRetries: 3,
		DialTimeout: 5*time.Second,
		ReadTimeout: 5*time.Second,
		WriteTimeout: 5*time.Second,
		MinIdleConns: minIdleConns,
	})
}

func GetRedis() *redis.Client{
	return RedisClient
}