package model

import (
	"time"

	"github.com/google/uuid"
)

type Products struct {
	ID        int        `gorm:"type:int"`
	UUID      uuid.UUID  `gorm:"type:varchar(100);uniqueIndex" json:"uuid"`
	Name      string     `gorm:"type:varchar(100)"`
	Image_URL string     `gorm:"type:varchar(255)"`
	Admin_ID  uint       `gorm:"type:int;ForeignKey:Admin_ID"`
	CreatedAt *time.Time `gorm:"type:timestamp"`
	UpdatedAt *time.Time `gorm:"type:timestamp"`
	Variants  []Variants `gorm:"ForeignKey:Product_ID;AssociationForeignKey:UUID"`
}
