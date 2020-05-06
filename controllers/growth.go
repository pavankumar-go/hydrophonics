package controllers

import (
	"fmt"

	"github.com/hydrophonics/database"
	"github.com/hydrophonics/models"
	"github.com/hydrophonics/types"
)

func AddGrowthInfo(info types.GrowthDTO) (*models.GrowthMonitor, error) {
	data := &models.GrowthMonitor{
		PlantID:   uint(info.PlantID),
		AirTemp:   info.AirTemp,
		WaterTemp: info.WaterTemp,
		Humidity:  info.Humidity,
		Color:     info.Color,
		Lighting:  info.Lighting,
	}

	var inf models.GrowthMonitor

	db := database.GetDB()
	err := db.Model(&inf).Create(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetGrowthInfoForPlant(plantID uint) ([]*models.GrowthMonitor, error) {
	fmt.Print(plantID)
	db := database.GetDB()

	data := &models.GrowthMonitor{PlantID: plantID}

	var info []*models.GrowthMonitor

	err := db.Where(data).Find(&info).Error
	if err != nil {
		return nil, err
	}

	return info, nil
}

// func GetPlant() (*[]models.Plant, error) {
// 	db := database.GetDB()
// 	var plant []models.Plant
// 	err := db.Where(&plant).Find(&plant).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &plant, nil
// }

// func DeletePlant(name string) (*models.Plant, error) {
// 	db := database.GetDB()

// 	data := &models.Plant{Name: name}

// 	var plant models.Plant

// 	err := db.Where(&plant).Unscoped().Delete(data).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &plant, nil
// }
