package main

import (
	"fmt"
	"net/http"

	"github.com/mhaatha/go-e-commerce-api/internal/bootstrap"
)

func main() {
	cfg, app := bootstrap.InitApp()

	http.ListenAndServe(fmt.Sprintf(":%s", cfg.AppPort), app)
}
