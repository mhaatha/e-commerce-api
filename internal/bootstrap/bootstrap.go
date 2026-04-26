package bootstrap

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/mhaatha/go-e-commerce-api/internal/config"
	"github.com/mhaatha/go-e-commerce-api/internal/controller"
	"github.com/mhaatha/go-e-commerce-api/internal/database"
	"github.com/mhaatha/go-e-commerce-api/internal/repository"
	"github.com/mhaatha/go-e-commerce-api/internal/service"
)

func InitApp() (*config.Config, http.Handler) {
	// Log init
	config.LogInit()

	// Config init
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("godotenv fails to load .env file", "err", err)
		os.Exit(1)
	}

	// Validator init
	validate := config.ValidatorInit()

	// Database init
	dbpool, err := database.ConnectDB(cfg)
	if err != nil {
		slog.Error("error has occured when calling ConnectDB function", "err", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	// Main router
	r := chi.NewRouter()

	// HealthStatus
	healthStatusController := controller.NewHealthStatusController()

	// Product
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, validate)
	productController := controller.NewProductController(productService)

	r.Get("/api/v1/health", healthStatusController.Check)

	r.Post("/api/v1/products", productController.Create)

	return cfg, r
}
