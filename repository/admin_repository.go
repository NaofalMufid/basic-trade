package repository

import (
	"basic-trade/model"

	"gorm.io/gorm"
)

type AdminRepository interface {
	Register(admin model.Admins)
	FindByEmail(email string) (*model.Admins, error)
}

type AdminRepositoryImpl struct {
	Db *gorm.DB
}

func NewAdminRepository(Db *gorm.DB) AdminRepository {
	return &AdminRepositoryImpl{Db: Db}
}

func (a AdminRepositoryImpl) Register(admin model.Admins) {
	result := a.Db.Create(&admin)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (a AdminRepositoryImpl) FindByEmail(email string) (*model.Admins, error) {
	var admin model.Admins

	if err := a.Db.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}
