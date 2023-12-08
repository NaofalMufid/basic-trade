package model

import (
	"time"
)

type Products struct {
	ID        int        `gorm:"type:int"`
	UUID      string     `gorm:"type:varchar(100);uniqueIndex" json:"uuid"`
	Name      string     `gorm:"type:varchar(100)"`
	Image_URL string     `gorm:"type:varchar(255)"`
	AdminID   uint       `gorm:"type:int;foreignKey:AdminID;references:ID"`
	CreatedAt *time.Time `gorm:"type:timestamp"`
	UpdatedAt *time.Time `gorm:"type:timestamp"`
	Variants  []Variants `gorm:"foreignKey:ProductID;references:UUID"`
}
