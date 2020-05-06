package types

type PlantDTO struct {
	Name           string `json:"name" binding:"required"`
	NumberOfPlants int    `json:"numberOfPlants" binding:"required"`
	CategoryID     uint   `json:"categoryID" binding:"required"`
}

type GrowthDTO struct {
	PlantID   int    `json:"plantID" binding:"required"`
	AirTemp   string `json:"airTemp" binding:"required"`
	WaterTemp string `json:"waterTemp" binding:"required"`
	Humidity  string `json:"humidity" binding:"required"`
	Color     string `json:"color" binding:"required"`
	Lighting  string `json:"lighting" binding:"required"`
}
