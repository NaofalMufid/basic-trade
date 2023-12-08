package service

import (
	"basic-trade/data/request"
	"basic-trade/data/response"
	"basic-trade/helper"
	"basic-trade/model"
	"basic-trade/repository"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type AdminService interface {
	Register(admin request.CreateAdminRequest) (response.AdminResponse, error)
	Login(email, password string) (response.LoginResponse, error)
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

func (a AdminServiceImpl) Register(admin request.CreateAdminRequest) (response.AdminResponse, error) {
	if err := a.Validate.Struct(admin); err != nil {
		return response.AdminResponse{}, fmt.Errorf("validation error: %v", err)
	}
	newUUID := uuid.New()
	adminModel := model.Admins{
		UUID:     newUUID,
		Name:     admin.Name,
		Email:    admin.Email,
		Password: admin.Password,
	}
	newAdmin, err := a.AdminRepository.Register(adminModel)
	if err != nil {
		return response.AdminResponse{}, err
	}
	return newAdmin, nil
}

func (a AdminServiceImpl) Login(email, password string) (response.LoginResponse, error) {
	admin, err := a.AdminRepository.FindByEmail(email)
	if err != nil {
		return response.LoginResponse{}, err
	}

	if !helper.ComparePassword([]byte(admin.Password), []byte(password)) {
		return response.LoginResponse{}, errors.New("invalid password")
	}
	token := helper.GenerateToken(admin.ID, admin.Email)
	data := response.LoginResponse{
		ID:    int(admin.ID),
		UUID:  admin.UUID,
		Name:  admin.Name,
		Email: admin.Email,
		Token: token,
	}
	return data, nil
}
