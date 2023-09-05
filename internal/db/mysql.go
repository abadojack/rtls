package db

import (
	"github.com/abadojack/rtls/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}

	db, err := gorm.Open(mysql.Open(config.AppConfig.DB), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
