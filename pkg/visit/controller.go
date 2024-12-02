package visit

import (
	"net/http"
    "github.com/go-chi/render"
	"clinic/config"
	"clinic/pkg/model"
	"clinic/database"
	"clinic/database/dbmodel"
)

type VisitConfig struct {
    *config.Config
}

func New(configuration *config.Config) *VisitConfig {
    return &VisitConfig{configuration}
}

func (config *VisitConfig) VisitHandler(w http.ResponseWriter, r *http.Request){
	req := &model.VisitRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error":"Invalid request payload"})
		return
	}

	visit := req.visit
	res := &model.VisitResponse{Visit: visit}
	render.JSON(w, r, res)

    visitEntry := dbmodel.VisitEntry{Visit: req.Visit}
    config.VisitEntryRepository.Create(&visitEntry)
}

