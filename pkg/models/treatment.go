package model

import (
	"net/http"
)

type TreatmentRequest struct {
	VisitID     int     `json:"visit_id"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
	Cost        float64 `json:"cost"`
}

func (a *TreatmentRequest) Bind(r *http.Request) error {
	return nil
}

type TreatmentResponse struct {
	VisitID     int     `json:"visit_id"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
	Cost        float64 `json:"cost"`
}
