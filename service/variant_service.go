package service

import "basic-trade/model"

type VariantService interface {
	GetAll() []model.Variants
	GetById(uuid string) (model.Variants, error)
	Create(variant model.Variants) error
	Update(uuid string, variant model.Variants) error
	Delete(uuid string) error
}
