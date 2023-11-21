package repository

import (
	"basic-trade/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() []model.Products
	FindByAdminID(adminID int) ([]model.Products, error)
	FindById(uuid string) (product model.Products, err error)
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

func (p ProductRepositoryImpl) FindAll() []model.Products {
	var products []model.Products
	result := p.Db.Find(&products)
	if result.Error != nil {
		panic(result.Error)
	}
	return products
}

func (p ProductRepositoryImpl) FindByAdminID(adminID int) ([]model.Products, error) {
	var products []model.Products
	if err := p.Db.Where("admin_id", adminID).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
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
