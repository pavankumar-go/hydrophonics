package controllers

import (
	"github.com/hydrophonics/database"
	"github.com/hydrophonics/models"
	"github.com/hydrophonics/types"
)

// AddPlant adds a plant
func AddPlant(plant types.PlantDTO) (*models.Plant, error) {
	data := &models.Plant{
		Name:       plant.Name,
		Count:      plant.NumberOfPlants,
		CategoryID: plant.CategoryID,
	}

	var p models.Plant
	db := database.GetDB()

	err := db.Model(&p).Create(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetPlants lists all plants
func GetPlants() (*[]models.Plant, error) {
	db := database.GetDB()
	var plant []models.Plant

	err := db.Where(&plant).Find(&plant).Error
	if err != nil {
		return nil, err
	}

	return &plant, nil
}

// GetPlant returns a plant thats matches the name
func GetPlant(name string) (*models.Plant, error) {
	data := &models.Plant{Name: name}
	db := database.GetDB()
	var plant models.Plant

	err := db.Where(&data).Find(&plant).Error
	if err != nil {
		return nil, err
	}

	return &plant, nil
}

// DeletePlant deletes a plant by its ID
func DeletePlant(id uint) (*models.Plant, error) {
	db := database.GetDB()
	data := &models.Plant{BaseModel: models.BaseModel{ID: id}}
	var plant models.Plant

	pID := &models.Plant{BaseModel: models.BaseModel{ID: id}}
	err := db.Where(&pID).Find(&plant).Error
	if err != nil {
		return nil, err
	}

	err = db.Where(&plant).Unscoped().Delete(data).Error
	if err != nil {
		return nil, err
	}

	return &plant, nil
}

// UpdatePlant updates a plant by its ID
func UpdatePlant(plantID uint, data interface{}) (*models.Plant, error) {
	db := database.GetDB()
	pID := &models.Plant{BaseModel: models.BaseModel{ID: plantID}}
	var plant models.Plant

	err := db.Where(&pID).Find(&plant).Error
	if err != nil {
		return nil, err
	}

	err = db.Model(&plant).Update(data).Error
	if err != nil {
		return nil, err
	}

	return &plant, nil
}
