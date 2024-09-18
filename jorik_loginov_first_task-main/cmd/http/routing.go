package httpserver

import (
	"encoding/json"
	"fmt"
	httphandlersv1 "github.com/akayo16/jorik_loginov_first_task/cmd/http/handlers/v1"
	"github.com/go-chi/chi/v5"
	"math/rand"
	"net/http"
	"strconv"
)

func GetRoutes() http.Handler {
	rootRouter := chi.NewRouter()

	rootRouter.Get("/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(
			fmt.Sprintf("Hello: %s", strconv.Itoa(rand.Int())),
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	rootRouter.Route("/api", func(r chi.Router) {
		r.Mount("/v1", httphandlersv1.ServiceV1())
	})

	return rootRouter
}
