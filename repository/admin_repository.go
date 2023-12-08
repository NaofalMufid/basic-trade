package repository

import (
	"basic-trade/data/response"
	"basic-trade/model"

	"gorm.io/gorm"
)

type AdminRepository interface {
	Register(admin model.Admins) (response.AdminResponse, error)
	FindByEmail(email string) (*model.Admins, error)
}

type AdminRepositoryImpl struct {
	Db *gorm.DB
}

func NewAdminRepository(Db *gorm.DB) AdminRepository {
	return &AdminRepositoryImpl{Db: Db}
}

func (a AdminRepositoryImpl) Register(admin model.Admins) (response.AdminResponse, error) {
	if err := admin.Validate(); err != nil {
		return response.AdminResponse{}, err
	}
	result := a.Db.Create(&admin)
	if result.Error != nil {
		return response.AdminResponse{}, result.Error
	}
	data := response.AdminResponse{
		ID:    int(admin.ID),
		UUID:  admin.UUID,
		Name:  admin.Name,
		Email: admin.Email,
	}
	return data, nil
}

func (a AdminRepositoryImpl) FindByEmail(email string) (*model.Admins, error) {
	var admin model.Admins

	if err := a.Db.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}
