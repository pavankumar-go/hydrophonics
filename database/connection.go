package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB can be accessible anywhere in project
var DB *gorm.DB

// Init establishes connection to database
func Init() *gorm.DB {
	// host := "localhost"
	// username := "postgres"
	// password := ""
	// dbName := "hydrophonicss"
	// port := 5432

	// databaseURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbName)

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=hydrophonics sslmode=disable")
	if err != nil {
		fmt.Println("Could not connect to database")
		return nil
	}

	db.DB().SetMaxIdleConns(10)
	db.LogMode(true)

	DB = db

	return DB
}

// GetDB re-initialises DB if nil
func GetDB() *gorm.DB {
	if DB == nil {
		Init()
	}
	return DB
}
