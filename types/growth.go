package types

// GrowthDTO adds growth info
type GrowthDTO struct {
	PlantID   int    `json:"plantID" binding:"required"`
	AirTemp   string `json:"airTemp" binding:"required"`
	WaterTemp string `json:"waterTemp" binding:"required"`
	Humidity  string `json:"humidity" binding:"required"`
	Color     string `json:"color" binding:"required"`
	Lighting  string `json:"lighting" binding:"required"`
}