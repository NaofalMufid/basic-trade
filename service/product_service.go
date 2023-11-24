package service

import (
	"basic-trade/helper"
	"basic-trade/model"
	"basic-trade/repository"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ProductService interface {
	GetAll() []model.Products
	GetByAdminID(adminID int) ([]model.Products, error)
	GetById(uuid string) model.Products
	Create(product model.Products) error
	Update(uuid string, product model.Products) error
	Delete(uuid string) error
}

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	Validate          *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		Validate:          validate,
	}
}

func (p ProductServiceImpl) Create(product model.Products) error {
	newUUID := uuid.New()
	productModel := model.Products{
		UUID:      newUUID,
		Name:      product.Name,
		Image_URL: product.Image_URL,
		Admin_ID:  product.Admin_ID,
	}
	p.ProductRepository.Save(productModel)
	return nil
}

func (p ProductServiceImpl) Update(uuid string, product model.Products) error {
	productData, err := p.ProductRepository.FindById(uuid)
	if err != nil {
		return err
	}

	if err := helper.DeleteFile(productData.Image_URL); err != nil {
		return err
	}

	productData.Name = product.Name
	productData.Image_URL = product.Image_URL
	err = p.ProductRepository.Update(productData)
	if err != nil {
		return err
	}
	return nil
}

func (p ProductServiceImpl) GetAll() []model.Products {
	result := p.ProductRepository.FindAll()
	var products []model.Products
	for _, v := range result {
		product := model.Products{
			ID:        v.ID,
			UUID:      v.UUID,
			Name:      v.Name,
			Image_URL: v.Image_URL,
			Admin_ID:  v.Admin_ID,
		}
		products = append(products, product)
	}
	return products
}

func (p ProductServiceImpl) GetByAdminID(adminID int) ([]model.Products, error) {
	result, err := p.ProductRepository.FindByAdminID(adminID)
	if err != nil {
		return nil, err
	}
	var products []model.Products
	for _, v := range result {
		product := model.Products{
			ID:        v.ID,
			UUID:      v.UUID,
			Name:      v.Name,
			Image_URL: v.Image_URL,
			Admin_ID:  v.Admin_ID,
		}
		products = append(products, product)
	}
	return products, nil
}

func (p ProductServiceImpl) GetById(uuid string) model.Products {
	product, err := p.ProductRepository.FindById(uuid)
	if err != nil {
		panic(err)
	}
	return product
}

func (p ProductServiceImpl) Delete(uuid string) error {
	productData, err := p.ProductRepository.FindById(uuid)
	if err != nil {
		return err
	}

	if err := helper.DeleteFile(productData.Image_URL); err != nil {
		return err
	}

	err = p.ProductRepository.Delete(uuid)
	if err != nil {
		return err
	}
	return nil
}
