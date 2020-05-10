package models

// GrowthInfo provides growth info for each type of plant
type GrowthInfo struct {
	BaseModel
	PlantID   uint   `gorm:"foreignkey:PlantID"`
	AirTemp   string `gorm:"varchar(10)"`
	WaterTemp string `gorm:"varchar(10)"`
	Humidity  string `gorm:"varchar(10)"`
	Color     string `gorm:"varchar(20)"`
	Lighting  string `gorm:"varchar(10)"`
}
