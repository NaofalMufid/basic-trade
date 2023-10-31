package model

import (
	"time"

	"github.com/google/uuid"
)

type Admins struct {
	ID        int        `gorm:"type:int"`
	UUID      uuid.UUID  `gorm:"type:varchar(255)"`
	Name      string     `gorm:"type:varchar(100)"`
	Email     string     `gorm:"type:varchar(100)"`
	Password  string     `gorm:"type:varchar(255)"`
	CreatedAt *time.Time `gorm:"type:timestamp"`
	UpdatedAt *time.Time `gorm:"type:timestamp"`
	Products  []Products `gorm:"ForeignKey:Admin_ID;AssociationForeignKey:UUID"`
}
