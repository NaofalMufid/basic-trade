package model

import (
	"time"

	"github.com/google/uuid"
)

type Variants struct {
	ID           int        `gorm:"type:int"`
	UUID         uuid.UUID  `gorm:"type:varchar(100);uniqueIndex" json:"uuid"`
	Variant_Name string     `gorm:"type:varchar(150)"`
	Quantity     int        `gorm:"type:int"`
	ProductID    uint       `gorm:"type:int;foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt    *time.Time `gorm:"type:timestamp"`
	UpdatedAt    *time.Time `gorm:"type:timestamp"`
}
