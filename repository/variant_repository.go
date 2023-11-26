package repository

import "basic-trade/model"

type VariantRepository interface {
	FindAll() ([]model.Variants, error)
	FindByID(uuid string) (model.Variants, error)
	Save(variant model.Variants) error
	Update(variant model.Variants) error
	Delete(uuid string) error
}
