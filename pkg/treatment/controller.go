package treatment

import (
	"clinic/config"
	"clinic/database/dbmodel"
	model "clinic/pkg/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type TreatmentConfig struct {
	*config.Config
}

func New(configuration *config.Config) *TreatmentConfig {
	return &TreatmentConfig{configuration}
}

func (config *TreatmentConfig) CreateTreatment(w http.ResponseWriter, r *http.Request) {
	req := &model.TreatmentRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	res := &model.TreatmentResponse{VisitID: req.VisitID, Description: req.Description, Date: req.Date, Cost: req.Cost}
	render.JSON(w, r, res)

	treatmentEntry := dbmodel.Treatment{VisitID: req.VisitID, Description: req.Description, Date: req.Date, Cost: req.Cost}
	config.TreatmentRepository.Create(&treatmentEntry)
}

func (config *TreatmentConfig) GetTreatments(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strId)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	treatment, err := config.TreatmentRepository.FindByVisit(id)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Error fetching treatment"})
		return
	}
	render.JSON(w, r, treatment)
}
