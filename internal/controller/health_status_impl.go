package controller

import (
	"net/http"
)

type HealthStatusImpl struct{}

func NewHealthStatusController() HealthStatusController {
	return &HealthStatusImpl{}
}

func (controller *HealthStatusImpl) Check(w http.ResponseWriter, r *http.Request) {
	// Not implemented yet
}
