package controller

import (
	"net/http"

	"github.com/mhaatha/go-e-commerce-api/internal/service"
)

type ProductControllerImpl struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		productService: productService,
	}
}

func (controller *ProductControllerImpl) Create(w http.ResponseWriter, r *http.Request) {}
