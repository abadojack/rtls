package models

import "github.com/sirupsen/logrus"

// GetLeaderBoard gets top n players
// if n players are not on redis, read from db and add them to redis
// else load from redis and return
func GetLeaderBoard(n int) ([]Player, error) {
	// if players not in redis cache
	if n > limit {
		players, err := getNPlayers(n)
		if err != nil {
			return nil, err
		}

		err = setLeaderboardToRedis(players)
		if err != nil {
			// log error but don't return it. Rather not exit when redis fails
			logrus.WithError(err).Error("error writing leaderboard to redis")
		}

		limit = len(players)
		last = players[limit-1]

		return players, nil
	}

	players, err := loadLeaderboardFromRedis()
	if err != nil {
		return nil, err
	}

	if len(players) > 0 {
		return players[:n], nil
	}

	return nil, nil
}
