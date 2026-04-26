package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/mhaatha/go-e-commerce-api/internal/model/web"
	"github.com/mhaatha/go-e-commerce-api/internal/repository"
)

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

func (service *ProductServiceImpl) CreateNewProduct(ctx context.Context, request web.CreateProductRequest) (web.CreateProductResponse, error) {
	return web.CreateProductResponse{}, nil
}
