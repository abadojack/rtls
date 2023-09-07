package models

import (
	"time"

	"github.com/abadojack/rtls/internal/db"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Player represents the player model
type Player struct {
	ID    string `json:"id"` //  By default, GORM uses ID as primary key
	Score int    `gorm:"index" json:"score"`
	Rank  int    `gorm:"-" json:"rank"` // rank is not stored in db

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// limit represents number of players saved in cache
var limit int

// last player on the leaderboard stored in redis
var last Player

// NewPlayer creates a new player to db with score
func NewPlayer(score int) (*Player, error) {
	db, err := db.GetDB()
	if err != nil {
		return nil, err
	}

	now := time.Now()

	player := &Player{
		ID:    uuid.NewString(),
		Score: score,

		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := db.Create(player).Error; err != nil {
		return nil, err
	}

	return player, nil
}

func (p *Player) Update() error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}

	return db.Model(&Player{}).Where("ID = ?", p.ID).Update("score", p.Score).Error
}

// Get a player from db including there rank
func (p *Player) Get() (*Player, error) {
	db, err := db.GetDB()
	if err != nil {
		return nil, err
	}

	type result struct {
		Count, Score int
	}

	var res result

	// Execute the SQL query
	err = db.Raw(`
	 SELECT 
		 (SELECT COUNT(*) FROM players WHERE score > c.score) AS count,
		 c.score
	 FROM 
		 players c
	 WHERE 
		 c.id = ?
    `, p.ID).Scan(&res).Error
	if err != nil {
		return nil, err
	}

	rank := res.Count + 1

	return &Player{
		ID:    p.ID,
		Score: res.Score,
		Rank:  rank,
	}, nil
}

// getNPlayers return top n players from db
func getNPlayers(n int) ([]Player, error) {
	db, err := db.GetDB()
	if err != nil {
		return nil, err
	}

	var players []Player

	err = db.Model(Player{}).Order("score DESC").Limit(n).Find(&players).Error
	if err != nil {
		return nil, err
	}

	noOfPlayers := len(players)

	if noOfPlayers == 0 {
		return nil, nil
	}

	// Calculate the rank for each player
	rank := 1
	prevScore := players[0].Score
	players[0].Rank = rank

	for i := 1; i < noOfPlayers; i++ {
		if players[i].Score == prevScore {
			players[i].Rank = rank
		} else {
			rank++
			players[i].Rank = rank
			prevScore = players[i].Score
		}
	}

	return players, nil
}

// if player is in top n, update top n
func (p *Player) updateCache() error {
	if p.Score >= last.Score {
		players, err := getNPlayers(limit)
		if err != nil {
			return err
		}

		logrus.Infoln("Updating redis cache")

		return setLeaderboardToRedis(players)
	}

	return nil
}

// AfterCreate is a hook that runs after create
func (p *Player) AfterCreate(tx *gorm.DB) error {
	return p.updateCache()
}

// AfterUpdate is a hook that runs after update
func (p *Player) AfterUpdate(tx *gorm.DB) error {
	return p.updateCache()
}
