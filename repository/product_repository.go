package repository

import (
	"basic-trade/data/response"
	"basic-trade/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(page, size int, search string) ([]model.Products, error)
	FindById(uuid string) (product model.Products, err error)
	CountProduct(search string) (int64, error)
	Save(product model.Products) (response.ProductResponse, error)
	Update(product model.Products) (response.ProductResponse, error)
	Delete(uuid string) error
}

type ProductRepositoryImpl struct {
	Db *gorm.DB
}

func NewProductRepository(Db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{Db: Db}
}

func (p ProductRepositoryImpl) Save(product model.Products) (response.ProductResponse, error) {
	result := p.Db.Create(&product)
	if result.Error != nil {
		return response.ProductResponse{}, result.Error
	}
	new_product := response.ProductResponse{
		ID:        product.ID,
		UUID:      product.UUID,
		Name:      product.Name,
		Image_URL: product.Image_URL,
		AdminID:   product.AdminID,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
	return new_product, nil
}

func (p ProductRepositoryImpl) Update(product model.Products) (response.ProductResponse, error) {
	if err := p.Db.Save(&product).Error; err != nil {
		return response.ProductResponse{}, err
	}
	update_product := response.ProductResponse{
		ID:        product.ID,
		UUID:      product.UUID,
		Name:      product.Name,
		Image_URL: product.Image_URL,
		AdminID:   product.AdminID,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
	return update_product, nil
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
	if err := p.Db.Where("uuid", uuid).Preload("Variants").First(&product).Error; err != nil {
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
