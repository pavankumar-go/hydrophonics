package models

// Plant contains all types of plants
type Plant struct {
	BaseModel
	Name       string `gorm:"varchar(50);unique_index"`
	Count      int    `gorm:"column:number_of_plants"`
	CategoryID uint   `gorm:"association_foreignkey:CategoryID"`
}
