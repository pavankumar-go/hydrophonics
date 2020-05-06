package models

// Category is where each plant is categorized
type Category struct {
	BaseModel
	Name string `gorm:"varchar(20); not null"`
}
