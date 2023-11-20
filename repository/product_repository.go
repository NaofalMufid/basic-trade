package repository

import (
	"basic-trade/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() []model.Products
	FindById(productId string) (product model.Products, err error)
	Save(product model.Products) error
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

func (p ProductRepositoryImpl) FindAll() []model.Products {
	var products []model.Products
	result := p.Db.Find(&products)
	if result.Error != nil {
		panic(result.Error)
	}
	return products
}

func (p ProductRepositoryImpl) FindById(productId string) (model.Products, error) {
	var product model.Products
	if err := p.Db.Where("uuid", productId).First(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}
