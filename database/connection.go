package database

import (
	"fmt"

	"github.com/hydrophonics/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/birkirb/loggers.v1/log"
)

// DB can be accessible anywhere in project
var DB *gorm.DB

// Init establishes connection to database
func Init() *gorm.DB {

	dbConnectionStr := GetDBAddress()
	db, err := gorm.Open("postgres", dbConnectionStr)
	if err != nil {
		log.Panic("Could not connect to database: ", err)
		return nil
	}

	db.DB().SetMaxIdleConns(10)
	db.LogMode(true)

	DB = db

	return DB
}

// GetDBAddress returns address string of database
func GetDBAddress() string {
	cfg := config.GetConfig()
	dbConnectionStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.DBName, cfg.Database.Password, cfg.Database.SSLMode)

	return dbConnectionStr
}

// GetDB re-initialises DB if nil
func GetDB() *gorm.DB {
	if DB == nil {
		Init()
	}
	return DB
}
