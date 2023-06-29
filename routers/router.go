package routers

import (
	"net/http"

	"github.com/FadyGamilM/cfe_api/apis"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func Routes() http.Handler {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		// ExposeHeaders:      []string{"Link"},
		// AllowedCredentials: true,
		MaxAge: 300,
	}))

	router.Get("/api/v1/coffees", apis.GetCoffees)

	return router
}
