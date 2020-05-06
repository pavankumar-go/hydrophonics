package migrate

import (
	"log"

	"github.com/hydrophonics/models"
	"github.com/jinzhu/gorm"
)

// Migrate creates tables and adds constraints
func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.Plant{}).Error
	if err != nil {
		log.Panicln("error migrating plants :", err)
	}

	err = db.AutoMigrate(&models.Category{}).Error
	if err != nil {
		log.Panicln("error migrating category :", err)
	}

	err = db.AutoMigrate(&models.GrowthMonitor{}).Error
	if err != nil {
		log.Panicln("error migrating plants :", err)
	}

	err = db.Model(&models.Plant{}).AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE").Error
	if err != nil {
		log.Panicln("error adding contraint to plants :", err)
	}

	err = db.Model(&models.GrowthMonitor{}).AddForeignKey("plant_id", "plants(id)", "CASCADE", "CASCADE").Error
	if err != nil {
		log.Panicln("error adding contraint to growth_monitor :", err)
	}
}
