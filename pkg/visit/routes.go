package visit

import (
    "github.com/go-chi/chi/v5"
	"clinic/config"
)

func Routes(configuration *config.Config) chi.Router {
	VisitConfig := New(configuration)
	router := chi.NewRouter()
	router.Post("/visit",VisitConfig.VisitHandler)
	return router
}