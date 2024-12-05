package cat

import (
	"clinic/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	CatConfig := New(configuration)
	router := chi.NewRouter()
	router.Post("/", CatConfig.CreateCat)
	router.Get("/", CatConfig.GetCats)
	router.Get("/{id}", CatConfig.GetCat)
	router.Put("/{id}", CatConfig.UpdateCat)
	router.Delete("/{id}", CatConfig.DeleteCat)
	return router
}
