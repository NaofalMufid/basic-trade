package service

import (
	"basic-trade/model"
	"basic-trade/repository"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ProductService interface {
	GetAll() []model.Products
	GetById(productId string) model.Products
	Create(product model.Products) error
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

func (p ProductServiceImpl) GetById(productId string) model.Products {
	product, err := p.ProductRepository.FindById(productId)
	if err != nil {
		panic(err)
	}
	return product
}
