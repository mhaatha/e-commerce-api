package repository

import (
	"context"

	"github.com/mhaatha/go-e-commerce-api/internal/model/domain"
)

type ProductRepository interface {
	Insert(ctx context.Context, product domain.Product) (domain.Product, error)
}
