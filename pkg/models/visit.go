package model

import (
	"net/http"
)

type VisitRequest struct {
	CatId  int       `json:"cat_id"`
	Date   time.Time `json:"date"`
	Reason string    `json:"reason"`
}

func (a *VisitRequest) Bind(r *http.Request) error {
    return nil
}

type VisitResponse struct {
	ID     int       `json:"id"`
	CatId  int       `json:"cat_id"`
	Date   time.Time `json:"date"`
	Reason string    `json:"reason"`
}