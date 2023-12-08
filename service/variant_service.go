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
	Create(variant model.Variants) (response.VariantResponse, error)
	Update(uuid string, variant model.Variants) (response.VariantResponse, error)
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
			ID:           v.ID,
			UUID:         v.UUID,
			Variant_Name: v.Variant_Name,
			Quantity:     v.Quantity,
			ProductID:    v.ProductID,
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

func (v VariantServiceImpl) Create(variant model.Variants) (response.VariantResponse, error) {
	newUUID := uuid.NewString()
	variantModel := model.Variants{
		UUID:         newUUID,
		Variant_Name: variant.Variant_Name,
		Quantity:     variant.Quantity,
		ProductID:    variant.ProductID,
	}
	new_variant, err := v.VariantRepository.Save(variantModel)
	if err != nil {
		return response.VariantResponse{}, err
	}
	return new_variant, nil
}

func (v VariantServiceImpl) Update(uuid string, variant model.Variants) (response.VariantResponse, error) {
	variantData, err := v.VariantRepository.FindByID(uuid)
	if err != nil {
		return response.VariantResponse{}, err
	}

	variantData.Variant_Name = variant.Variant_Name
	variantData.Quantity = variant.Quantity
	variantUPdate, err := v.VariantRepository.Update(variantData)
	if err != nil {
		return response.VariantResponse{}, err
	}
	return variantUPdate, nil
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
