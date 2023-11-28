package service

import (
	"basic-trade/data/response"
	"basic-trade/helper"
	"basic-trade/model"
	"basic-trade/repository"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ProductService interface {
	GetAll(page, size int, search string) (response.PaginatedProductResponse, error)
	GetById(uuid string) (response.ProductResponse, error)
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
		AdminID:   product.AdminID,
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

func (p ProductServiceImpl) GetAll(page, size int, search string) (response.PaginatedProductResponse, error) {
	paginator := helper.NewPagination(page, size)

	result, err := p.ProductRepository.FindAll(paginator.Page, paginator.PageSize, search)
	if err != nil {
		return response.PaginatedProductResponse{}, err
	}

	var products []response.ProductResponse
	for _, v := range result {
		product := response.ProductResponse{
			ID:        v.ID,
			UUID:      v.UUID,
			Name:      v.Name,
			Image_URL: v.Image_URL,
			AdminID:   v.AdminID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		for _, variant := range v.Variants {
			variantData := response.VariantResponse{
				ID:           variant.ID,
				UUID:         variant.UUID,
				Variant_Name: variant.Variant_Name,
				ProductID:    variant.ProductID,
				Quantity:     variant.Quantity,
				CreatedAt:    variant.CreatedAt,
				UpdatedAt:    variant.UpdatedAt,
			}
			product.Variants = append(product.Variants, variantData)
		}
		products = append(products, product)
	}

	totalProduct, err := p.ProductRepository.CountProduct(search)
	if err != nil {
		return response.PaginatedProductResponse{}, err
	}

	totalPage := paginator.TotalPage(totalProduct)

	productResponse := response.PaginatedProductResponse{
		Page:      paginator.Page,
		PageSize:  paginator.PageSize,
		TotalPage: totalPage,
		TotalData: totalProduct,
		Data:      products,
	}

	return productResponse, nil
}

func (p ProductServiceImpl) GetById(uuid string) (response.ProductResponse, error) {
	product, err := p.ProductRepository.FindById(uuid)
	if err != nil {
		panic(err)
	}
	data := response.ProductResponse{
		ID:        product.ID,
		Name:      product.Name,
		Image_URL: product.Image_URL,
		AdminID:   product.AdminID,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
	for _, v := range product.Variants {
		variant := response.VariantResponse{
			ID:           v.ID,
			UUID:         v.UUID,
			Variant_Name: v.Variant_Name,
			ProductID:    v.ProductID,
			Quantity:     v.Quantity,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		}
		data.Variants = append(data.Variants, variant)
	}
	return data, nil
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
