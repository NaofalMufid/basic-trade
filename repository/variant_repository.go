package repository

import (
	"basic-trade/model"

	"gorm.io/gorm"
)

type VariantRepository interface {
	FindAll(page, size int, search string) ([]model.Variants, error)
	FindByID(uuid string) (model.Variants, error)
	CountVariant(search string) (int64, error)
	Save(variant model.Variants) error
	Update(variant model.Variants) error
	Delete(uuid string) error
}

type VariantRepositoryImpl struct {
	Db *gorm.DB
}

func NewVariantRepository(Db *gorm.DB) VariantRepository {
	return &VariantRepositoryImpl{Db: Db}
}

func (v VariantRepositoryImpl) FindAll(page, size int, search string) ([]model.Variants, error) {
	var variants []model.Variants
	result := v.Db.Model(&model.Variants{})
	if search != "" {
		result = result.Where("variant_name LIKE ?", "%"+search+"%")
	}
	if size > 0 {
		offset := (page - 1) * size
		result = result.Offset(offset).Limit(size)
	}

	if err := result.Find(&variants).Error; err != nil {
		return nil, err
	}
	return variants, nil
}

func (v VariantRepositoryImpl) CountVariant(search string) (int64, error) {
	var count int64
	query := v.Db.Model(&model.Variants{})

	if search != "" {
		query = query.Where("variant_name LIKE ?", "%"+search+"%")
	}
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
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
