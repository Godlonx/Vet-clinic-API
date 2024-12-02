package model

import (
	"time"
	"net/http"
)

type TreatmentRequest struct {
	PetID       int       `json:"pet_id"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Cost        float64   `json:"cost"`
}

func (a *TreatmentRequest) Bind(r *http.Request) error {
    return nil
}

type TreatmentResponse struct {
	ID          int       `json:"id"`
	PetID       int       `json:"pet_id"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Cost        float64   `json:"cost"`
}