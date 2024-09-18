package httphandlersv1

import (
	httphandlersv1product "github.com/akayo16/jorik_loginov_first_task/cmd/http/handlers/v1/product"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func ServiceV1() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	// TODO More Detailed Cors Customization
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowOriginFunc:    nil,
		AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:     []string{"Link"},
		AllowCredentials:   false,
		MaxAge:             300,
		OptionsPassthrough: false,
		Debug:              false,
	}))

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.DocExpansion("docs/transfer/http/v1/swagger"), //The url pointing to API definition
	))

	r.Mount("/product", httphandlersv1product.NewProductRouter())

	return r
}
