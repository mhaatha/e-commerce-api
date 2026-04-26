package service

import (
	"context"

	"github.com/mhaatha/go-e-commerce-api/internal/model/web"
)

type ProductService interface {
	CreateNewProduct(ctx context.Context, request web.CreateProductRequest) (web.CreateProductResponse, error)
}
