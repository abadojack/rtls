package db

import (
	"github.com/abadojack/rtls/config"
	"github.com/redis/go-redis/v9"
)

const (
	LeaderboardHash = "leaderboard"
	LeaderboardKey  = "players"
)

var redisClient *redis.Client

func GetRedis() *redis.Client {
	if redisClient != nil {
		return redisClient
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.RedisHost,
		Password: config.AppConfig.RedisPassword,
		DB:       0,
	})

	return redisClient
}
