package repository

import (
	"basic-trade/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(page, size int, search string) ([]model.Products, error)
	FindById(uuid string) (product model.Products, err error)
	CountProduct(search string) (int64, error)
	Save(product model.Products) error
	Update(product model.Products) error
	Delete(uuid string) error
}

type ProductRepositoryImpl struct {
	Db *gorm.DB
}

func NewProductRepository(Db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{Db: Db}
}

func (p ProductRepositoryImpl) Save(product model.Products) error {
	result := p.Db.Create(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p ProductRepositoryImpl) Update(product model.Products) error {
	if err := p.Db.Save(&product).Error; err != nil {
		return err
	}
	return nil
}

func (p ProductRepositoryImpl) FindAll(page, size int, search string) ([]model.Products, error) {
	var products []model.Products
	query := p.Db.Model(&model.Products{}).Preload("Variants")

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	if size > 0 {
		offset := (page - 1) * size
		query = query.Offset(offset).Limit(size)
	}

	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p ProductRepositoryImpl) CountProduct(search string) (int64, error) {
	var count int64
	query := p.Db.Model(&model.Products{})

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (p ProductRepositoryImpl) FindById(uuid string) (model.Products, error) {
	var product model.Products
	if err := p.Db.Where("uuid", uuid).First(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (p ProductRepositoryImpl) Delete(uuid string) error {
	var product model.Products
	if err := p.Db.Where("uuid", uuid).Delete(&product).Error; err != nil {
		return err
	}
	return nil
}
