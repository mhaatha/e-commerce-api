package repository

import (
	"context"

	"github.com/mhaatha/go-e-commerce-api/internal/model/domain"
)

type ProductRepositoryImpl struct{}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Insert(ctx context.Context, product domain.Product) (domain.Product, error) {
	return domain.Product{}, nil
}
