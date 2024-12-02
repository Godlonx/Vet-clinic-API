package treatment

import (
	"net/http"
    "github.com/go-chi/render"
	"clinic/config"
	"clinic/pkg/models"
	"clinic/database/dbmodel"
)

type TreatmentConfig struct {
    *config.Config
}

func New(configuration *config.Config) *TreatmentConfig {
    return &TreatmentConfig{configuration}
}

func (config *TreatmentConfig) TreatmentHandler(w http.ResponseWriter, r *http.Request){
	req := &model.TreatmentRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error":"Invalid request payload"})
		return
	}

	res := &model.TreatmentResponse{VisitID: req.VisitID, Description: req.Description, Date: req.Date, Cost: req.Cost}
	render.JSON(w, r, res)

    treatmentEntry := dbmodel.Treatment{VisitID: req.VisitID, Description: req.Description, Date: req.Date, Cost: req.Cost}
    config.TreatmentRepository.Create(&treatmentEntry)
}

