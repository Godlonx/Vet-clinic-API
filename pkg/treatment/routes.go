package treatment

import (
    "github.com/go-chi/chi/v5"
	"clinic/config"
)

func Routes(configuration *config.Config) chi.Router {
	TreatmentConfig := New(configuration)
	router := chi.NewRouter()
	router.Post("/treatment",TreatmentConfig.TreatmentHandler)
	return router
}