package controller

import "net/http"

type HealthStatusController interface {
	Check(w http.ResponseWriter, r *http.Request)
}
