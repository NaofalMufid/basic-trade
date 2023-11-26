package service

import (
	"basic-trade/model"
	"basic-trade/repository"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

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

func (v VariantServiceImpl) GetAll() []model.Variants {
	result, _ := v.VariantRepository.FindAll()
	var variants []model.Variants
	for _, v := range result {
		variant := model.Variants{
			ID:           v.ID,
			UUID:         v.UUID,
			Variant_Name: v.Variant_Name,
			Quantity:     v.Quantity,
			Product_ID:   v.Product_ID,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		}
		variants = append(variants, variant)
	}
	return variants
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
