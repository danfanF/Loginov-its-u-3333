package httphandlersv1product

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewProductRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/get/all", All)

	return r
}
