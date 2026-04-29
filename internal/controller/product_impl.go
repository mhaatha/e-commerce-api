package controller

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/mhaatha/go-e-commerce-api/internal/model/web"
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

func (controller *ProductControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	var bodyBytes []byte
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("error when read the request body", "err", err)
	}

	var productRequest web.CreateProductRequest
	if err := json.Unmarshal(bodyBytes, &productRequest); err != nil {
		slog.Error("error when unmarshal the JSON", "err", err)
	}

	productResponse, err := controller.productService.CreateNewProduct(r.Context(), productRequest)
	if err != nil {
		slog.Error("error when calling the create new product service", "err", err)
	}

	jsonResponse, err := json.Marshal(productResponse)
	if err != nil {
		slog.Error("error when marshal the JSON", "err", err)
	}
}
