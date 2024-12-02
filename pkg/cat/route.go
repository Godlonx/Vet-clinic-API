package cat

import (
    "github.com/go-chi/chi/v5"
)

func Routes() chi.Router {
	router := chi.NewRouter()
	router.Post("/cat",CatHandler)
	router.Get("/history", CatHistoryHandler)
	return router
}