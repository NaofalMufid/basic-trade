package repository

import (
	"basic-trade/model"

	"gorm.io/gorm"
)

type VariantRepositoryImpl struct {
	Db *gorm.DB
}

func NewVariantRepository(Db *gorm.DB) VariantRepository {
	return &VariantRepositoryImpl{Db: Db}
}

func (v VariantRepositoryImpl) FindAll() ([]model.Variants, error) {
	var variants []model.Variants
	result := v.Db.Find(&variants)
	if result.Error != nil {
		return variants, result.Error
	}
	return variants, nil
}

func (v VariantRepositoryImpl) FindByID(uuid string) (model.Variants, error) {
	var variant model.Variants
	if err := v.Db.Where("uuid", uuid).First(&variant).Error; err != nil {
		return variant, err
	}
	return variant, nil
}

func (v VariantRepositoryImpl) Save(variant model.Variants) error {
	result := v.Db.Create(&variant)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (v VariantRepositoryImpl) Update(variant model.Variants) error {
	if err := v.Db.Save(&variant).Error; err != nil {
		return err
	}
	return nil
}

func (v VariantRepositoryImpl) Delete(uuid string) error {
	var variant model.Variants
	if err := v.Db.Where("uuid", uuid).Delete(&variant).Error; err != nil {
		return err
	}
	return nil
}
