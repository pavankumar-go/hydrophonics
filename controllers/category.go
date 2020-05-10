package controllers

import (
	"github.com/hydrophonics/database"
	"github.com/hydrophonics/models"
)

// AddCategory adds 3 categories of plant
func AddCategory(name string) (*models.Category, error) {
	data := &models.Category{
		Name: name,
	}

	var category models.Category
	var count int

	db := database.GetDB()
	err := db.Where(data).First(&category).Count(&count).Error
	if err != nil {
		if err.Error() == "record not found" {
			if count == 0 {
				err = db.Model(&category).Create(data).Error
				if err != nil {
					return nil, err
				}
			}
			return nil, nil
		}
		return nil, err
	}
	return &category, nil
}

// GetCategories lists all categories
func GetCategories() (*[]models.Category, error) {
	db := database.GetDB()
	var category []models.Category

	err := db.Find(&category).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}
