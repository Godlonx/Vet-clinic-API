package main

import (
	"clinic/config"
	"clinic/pkg/cat"
	"clinic/pkg/treatment"
	"clinic/pkg/visit"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	configuration, err := config.New()
	if err != nil {
		log.Panicln("Configuration error:", err)
	}

	// Initialisation des routes
	router := Routes(configuration)

	log.Println("Serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Mount("/api/v1/cats", cat.Routes(configuration))
	router.Mount("/api/v1/visits", visit.Routes(configuration))
	router.Mount("/api/v1/treatments", treatment.Routes(configuration))
	return router
}
