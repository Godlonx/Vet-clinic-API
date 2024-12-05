package visit

import (
	"clinic/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	VisitConfig := New(configuration)
	router := chi.NewRouter()
	router.Post("/", VisitConfig.CreateVisit)
	router.Get("/{id}", VisitConfig.GetVisits)
	return router
}
