package models

import (
	"time"

	"github.com/abadojack/rtls/internal/db"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Player struct {
	ID    string //  By default, GORM uses ID as primary key
	Score int    `gorm:"index"`
	Rank  int    `gorm:"-"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

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

	return db.Where("ID = ?", p.ID).Update("score", p.Score).Error
}

func (p *Player) Get() (*Player, error) {
	db, err := db.GetDB()
	if err != nil {
		return nil, err
	}

	type result struct {
		count, score int
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

	rank := res.count + 1

	return &Player{
		ID:    p.ID,
		Score: res.score,
		Rank:  rank,
	}, nil
}
