package db

import (
	"fmt"

	"github.com/abadojack/rtls/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}

	dbUser := config.AppConfig.DBUser
	dbPass := config.AppConfig.DBPassword
	dbHost := config.AppConfig.DBHost
	dbName := config.AppConfig.DBName

	// Create the database connection string
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
