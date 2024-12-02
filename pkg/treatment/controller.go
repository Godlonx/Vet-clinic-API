package treatment

import (
	"net/http"
    "github.com/go-chi/render"
	"clinic/config"
	"clinic/pkg/model"
	"clinic/database"
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

	treatment := req.Treatment
	res := &model.TreatmentResponse{Treatment: treatment}
	render.JSON(w, r, res)

    treatmentEntry := dbmodel.TreatmentEntry{Treatment: req.Treatment}
    config.TreatmentEntryRepository.Create(&treatmentEntry)
}

