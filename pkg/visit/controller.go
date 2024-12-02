package visit

import (
	"net/http"
    "github.com/go-chi/render"
	"clinic/config"
	"clinic/pkg/models"
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

	res := &model.VisitResponse{CatId: req.CatId, Date: req.Date, Reason: req.Reason,CareTakerId: req.CareTakerId}
	render.JSON(w, r, res)

    visitEntry := dbmodel.Visit{CatId: req.CatId, Date: req.Date, Reason: req.Reason, CareTakerId: req.CareTakerId}
    err := config.VisitRepository.Create(&visitEntry)

	if err != nil {
		render.JSON(w, r, map[string]string{"error":"Error creating visit"})
		return
	}
}

