package models

import (
	"time"
)

// BaseModel defines the basic attributes that every other entity needs to implement
type BaseModel struct {
	ID        uint `gorm:"primary_key; not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
