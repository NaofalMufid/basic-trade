package model

import (
	"basic-trade/helper"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admins struct {
	ID        uint       `gorm:"type:int" json:"id"`
	UUID      uuid.UUID  `gorm:"type:varchar(100);uniqueIndex" json:"uuid"`
	Name      string     `gorm:"type:varchar(100)" json:"name"`
	Email     string     `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password  string     `gorm:"type:varchar(255);not null" json:"password"`
	CreatedAt *time.Time `gorm:"type:timestamp" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"type:timestamp" json:"updated_at,omitempty"`
	Products  []Products `gorm:"ForeignKey:AdminID;AssociationForeignKey:UUID" json:"products"`
}

func (a *Admins) BeforeCreate(tx *gorm.DB) (err error) {
	if a.Password != "" {
		a.Password = helper.HashPassword(a.Password)
	}
	err = nil
	return
}

func (a *Admins) Validate() error {
	validate := validator.New()
	return validate.Struct(a)
}
