package controllers

import (
	"errors"
	"fmt"

	"github.com/hydrophonics/database"
	"github.com/hydrophonics/models"
	"github.com/hydrophonics/types"
	"github.com/jinzhu/gorm"
	"gopkg.in/birkirb/loggers.v1/log"
)

// AddGrowthInfo addsngrowth info for a plant
func AddGrowthInfo(info types.GrowthDTO) (*models.GrowthInfo, error) {
	data := &models.GrowthInfo{
		PlantID:   uint(info.PlantID),
		AirTemp:   info.AirTemp,
		WaterTemp: info.WaterTemp,
		Humidity:  info.Humidity,
		Color:     info.Color,
		Lighting:  info.Lighting,
	}
	plantData := &models.Plant{BaseModel: models.BaseModel{ID: uint(info.PlantID)}}

	var growthInfo models.GrowthInfo
	var plant models.Plant

	db := database.GetDB()
	var count int
	err := db.Where(plantData).Find(&plant).Count(&count).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, errors.New("plant does not exist")
		}
		return nil, err
	}
	log.Println(count)

	if count > 1 {
		return nil, errors.New("growth info already exists for a plant")
	}

	err = db.Model(&growthInfo).Create(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetGrowthInfoForPlant lists growth-info added for a plant
func GetGrowthInfoForPlant(plantID uint) ([]*models.GrowthInfo, error) {
	fmt.Print(plantID)
	db := database.GetDB()

	data := &models.GrowthInfo{PlantID: plantID}

	var info []*models.GrowthInfo

	err := db.Where(data).Find(&info).Error
	if err != nil {
		return nil, err
	}

	return info, nil
}

// DeleteGrowthInfoOfPlant deletes growth-info of a plant
func DeleteGrowthInfoOfPlant(plantID uint, growthInfoID uint) (*models.GrowthInfo, error) {
	db := database.GetDB()

	data := &models.GrowthInfo{BaseModel: models.BaseModel{ID: growthInfoID}, PlantID: plantID}

	var growthInfo models.GrowthInfo

	err := db.Where(data).Find(&growthInfo).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, errors.New("growth info not found")
		}
		return nil, err
	}

	err = db.Where(data).Unscoped().Delete(&growthInfo).Error
	if err != nil {
		return nil, err
	}
	return &growthInfo, nil
}
