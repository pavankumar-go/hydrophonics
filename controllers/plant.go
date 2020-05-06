package controllers

import (
	"github.com/hydrophonics/database"
	"github.com/hydrophonics/models"
	"github.com/hydrophonics/types"
)

func AddPlants(plant types.PlantDTO) (*models.Plant, error) {
	data := &models.Plant{
		Name:       plant.Name,
		Count:      plant.NumberOfPlants,
		CategoryID: plant.CategoryID,
	}

	var pl models.Plant

	db := database.GetDB()
	err := db.Model(&pl).Create(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetPlants() (*[]models.Plant, error) {
	db := database.GetDB()
	var plant []models.Plant
	err := db.Where(&plant).Find(&plant).Error
	if err != nil {
		return nil, err
	}
	return &plant, nil
}

func GetPlant() (*[]models.Plant, error) {
	db := database.GetDB()
	var plant []models.Plant
	err := db.Where(&plant).Find(&plant).Error
	if err != nil {
		return nil, err
	}
	return &plant, nil
}

func DeletePlant(id uint) (*models.Plant, error) {
	db := database.GetDB()

	data := &models.Plant{BaseModel: models.BaseModel{ID: id}}

	var plant models.Plant

	err := db.Where(&plant).Unscoped().Delete(data).Error
	if err != nil {
		return nil, err
	}
	return &plant, nil
}
