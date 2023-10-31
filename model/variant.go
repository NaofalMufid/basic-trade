package model

import (
	"time"

	"github.com/google/uuid"
)

type Variants struct {
	ID           int        `gorm:"type:int"`
	UUID         uuid.UUID  `gorm:"type:varchar(255)"`
	Variant_Name string     `gorm:"type:varchar(150)"`
	Quantity     string     `gorm:"type:int"`
	Product_ID   string     `gorm:"type:varchar(255);ForeignKey:Product_ID"`
	CreatedAt    *time.Time `gorm:"type:timestamp"`
	UpdatedAt    *time.Time `gorm:"type:timestamp"`
}
