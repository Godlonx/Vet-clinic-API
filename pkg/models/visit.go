package model

import (
	"net/http"
)

type VisitRequest struct {
	CatId     int    `json:"cat_id"`
	Date      string `json:"date"`
	Reason    string `json:"reason"`
	CareTaker string `json:"care_taker"`
}

func (a *VisitRequest) Bind(r *http.Request) error {
	return nil
}

type VisitResponse struct {
	CatId     int    `json:"cat_id"`
	Date      string `json:"date"`
	Reason    string `json:"reason"`
	CareTaker string `json:"care_taker"`
}
