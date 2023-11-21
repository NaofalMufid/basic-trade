package service

import (
	"basic-trade/data/request"
	"basic-trade/helper"
	"basic-trade/model"
	"basic-trade/repository"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type AdminService interface {
	Register(admin request.CreateAdminRequest) error
	Login(email, password string) (string, error)
}

type AdminServiceImpl struct {
	AdminRepository repository.AdminRepository
	Validate        *validator.Validate
}

func NewAdminServiceImpl(adminRepository repository.AdminRepository, validate *validator.Validate) AdminService {
	return &AdminServiceImpl{
		AdminRepository: adminRepository,
		Validate:        validate,
	}
}

func (a AdminServiceImpl) Register(admin request.CreateAdminRequest) error {
	if err := a.Validate.Struct(admin); err != nil {
		return err
	}
	newUUID := uuid.New()
	adminModel := model.Admins{
		UUID:     newUUID,
		Name:     admin.Name,
		Email:    admin.Email,
		Password: admin.Password,
	}
	a.AdminRepository.Register(adminModel)
	return nil
}

func (a AdminServiceImpl) Login(email, password string) (string, error) {
	admin, err := a.AdminRepository.FindByEmail(email)
	if err != nil {
		return "", err
	}

	if !helper.ComparePassword([]byte(admin.Password), []byte(password)) {
		return "", errors.New("Invalid Password")
	}

	token := helper.GenerateToken(admin.ID, admin.Email)
	return token, nil
}