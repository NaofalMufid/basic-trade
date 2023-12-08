package model

import (
	"time"
)

type Variants struct {
	ID           int        `gorm:"type:int"`
	UUID         string     `gorm:"type:varchar(100);uniqueIndex" json:"uuid"`
	Variant_Name string     `gorm:"type:varchar(150)"`
	Quantity     int        `gorm:"type:int"`
	ProductID    string     `gorm:"type:varchar(100);foreignKey:ProductID;references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt    *time.Time `gorm:"type:timestamp"`
	UpdatedAt    *time.Time `gorm:"type:timestamp"`
}
