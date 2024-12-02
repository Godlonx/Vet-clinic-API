package cat

import (
    "github.com/go-chi/chi/v5"
	"clinic/config"
)

func Routes(configuration *config.Config) chi.Router {
	CatConfig := New(configuration)
	router := chi.NewRouter()
	router.Post("/cat",CatConfig.CatHandler)
	return router
}