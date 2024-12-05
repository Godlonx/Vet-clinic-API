package treatment

import (
	"clinic/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	TreatmentConfig := New(configuration)
	router := chi.NewRouter()
	router.Post("/", TreatmentConfig.CreateTreatment)
	router.Get("/{id}", TreatmentConfig.GetTreatments)
	return router
}
