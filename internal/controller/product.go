package controller

import "net/http"

type ProductController interface {
	Create(w http.ResponseWriter, r *http.Request)
}
