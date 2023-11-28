package service

import (
	"basic-trade/data/response"
	"basic-trade/helper"
	"basic-trade/model"
	"basic-trade/repository"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type VariantService interface {
	GetAll(page, size int, search string) (response.PaginatedVariantResponse, error)
	GetById(uuid string) (model.Variants, error)
	Create(variant model.Variants) error
	Update(uuid string, variant model.Variants) error
	Delete(uuid string) error
}

type VariantServiceImpl struct {
	VariantRepository repository.VariantRepository
	Validate          *validator.Validate
}

func NewVariantService(variantRepository repository.VariantRepository, validate *validator.Validate) VariantService {
	return &VariantServiceImpl{
		VariantRepository: variantRepository,
		Validate:          validate,
	}
}

func (v VariantServiceImpl) GetAll(page, size int, search string) (response.PaginatedVariantResponse, error) {
	paginator := helper.NewPagination(page, size)

	result, err := v.VariantRepository.FindAll(paginator.Page, paginator.PageSize, search)
	if err != nil {
		return response.PaginatedVariantResponse{}, err
	}

	var variants []response.VariantResponse
	for _, v := range result {
		variant := response.VariantResponse{
			UUID:         v.UUID,
			Variant_Name: v.Variant_Name,
			Quantity:     v.Quantity,
			Product_ID:   v.Product_ID,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		}
		variants = append(variants, variant)
	}

	totalVariant, err := v.VariantRepository.CountVariant(search)
	if err != nil {
		return response.PaginatedVariantResponse{}, err
	}

	totalPage := paginator.TotalPage(totalVariant)

	variantResponse := response.PaginatedVariantResponse{
		Page:      paginator.Page,
		PageSize:  paginator.PageSize,
		TotalPage: totalPage,
		TotalData: totalVariant,
		Data:      variants,
	}

	return variantResponse, nil
}

func (v VariantServiceImpl) GetById(uuid string) (model.Variants, error) {
	variant, err := v.VariantRepository.FindByID(uuid)
	if err != nil {
		return variant, err
	}
	return variant, nil
}

func (v VariantServiceImpl) Create(variant model.Variants) error {
	newUUID := uuid.New()
	variantModel := model.Variants{
		UUID:         newUUID,
		Variant_Name: variant.Variant_Name,
		Quantity:     variant.Quantity,
		Product_ID:   variant.Product_ID,
	}
	v.VariantRepository.Save(variantModel)
	return nil
}

func (v VariantServiceImpl) Update(uuid string, variant model.Variants) error {
	variantData, err := v.VariantRepository.FindByID(uuid)
	if err != nil {
		return err
	}

	variantData.Variant_Name = variant.Variant_Name
	variantData.Quantity = variant.Quantity
	err = v.VariantRepository.Update(variantData)
	if err != nil {
		return err
	}
	return nil
}

func (v VariantServiceImpl) Delete(uuid string) error {
	_, err := v.VariantRepository.FindByID(uuid)
	if err != nil {
		return err
	}

	err = v.VariantRepository.Delete(uuid)
	if err != nil {
		return err
	}
	return nil
}