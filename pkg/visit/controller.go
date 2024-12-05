package visit

import (
	"clinic/config"
	"clinic/database/dbmodel"
	model "clinic/pkg/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type VisitConfig struct {
	*config.Config
}

func New(configuration *config.Config) *VisitConfig {
	return &VisitConfig{configuration}
}

func (config *VisitConfig) CreateVisit(w http.ResponseWriter, r *http.Request) {
	req := &model.VisitRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	res := &model.VisitResponse{CatId: req.CatId, Date: req.Date, Reason: req.Reason, CareTaker: req.CareTaker}
	render.JSON(w, r, res)

	visitEntry := dbmodel.Visit{CatId: req.CatId, Date: req.Date, Reason: req.Reason, CareTaker: req.CareTaker}
	err := config.VisitRepository.Create(&visitEntry)

	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Error creating visit"})
		return
	}
}

func (config *VisitConfig) GetVisits(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	catId, err := strconv.Atoi(strId)
	if err != nil || catId < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	visits, err := config.VisitRepository.FindAllByCatId(catId)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Error fetching visits"})
		return
	}
	render.JSON(w, r, visits)
}
