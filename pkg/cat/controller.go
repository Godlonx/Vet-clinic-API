package cat

import (
	"net/http"
    "github.com/go-chi/render"
	"clinic/config"
	"clinic/pkg/model"
	"clinic/database"
	"clinic/database/dbmodel"
)

type CatConfig struct {
    *config.Config
}

func New(configuration *config.Config) *CatConfig {
    return &CatConfig{configuration}
}

func (config *CatConfig) CatHandler(w http.ResponseWriter, r *http.Request){
	req := &model.CatRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error":"Invalid request payload"})
		return
	}

	cat := req.Cat
	res := &model.CatResponse{Cat: cat}
	render.JSON(w, r, res)

	catEntry := dbmodel.CatEntry{Cat: req.Cat}
    config.CatEntryRepository.Create(&catEntry)

}

