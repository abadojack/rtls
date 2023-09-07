package models

import (
	"context"
	"encoding/json"

	"github.com/abadojack/rtls/internal/db"
)

// setLeaderboardToRedis write player leaderboard to redis
func setLeaderboardToRedis(leaderboard []Player) error {
	// Serialize the struct to JSON
	jsonData, err := json.Marshal(leaderboard)
	if err != nil {
		return err
	}

	return db.GetRedis().HSet(context.Background(), db.LeaderboardHash, db.LeaderboardKey, jsonData).Err()
}

// loadLeaderboardToRedis reads player leaderboard from redis
func loadLeaderboardFromRedis() ([]Player, error) {
	var leaderboard []Player

	result, _ := db.GetRedis().HGetAll(context.Background(), db.LeaderboardHash).Result()

	err := json.Unmarshal([]byte(result["players"]), &leaderboard)
	if err != nil {
		return nil, err
	}

	return leaderboard, nil
}
