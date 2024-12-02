package model

import (
	"net/http"
)

type CatRequest struct {
    Age   int    `json:"age"`
	Breed string `json:"breed"`
	Name  string `json:"name"`
	Weight float64 `json:"weight"`
}

func (a *CatRequest) Bind(r *http.Request) error {
    return nil
}

type CatResponse struct {
	ID	int    `json:"id"`
	Age   int    `json:"age"`
	Breed string `json:"breed"`
	Name  string `json:"name"`
	Weight float64 `json:"weight"`
}