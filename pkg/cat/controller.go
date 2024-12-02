package cat

import (
	"net/http"
    "github.com/go-chi/render"
	"vet-clinic-api/pkg/model"
	"vet-clinic-api/database"
	"vet-clinic-api/database/dbmodel"
)

func CatHandler(w http.ResponseWriter, r *http.Request){
	req := &model.CatRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error":"Invalid request payload"})
		return
	}

	cat := req.Cat
	res := &model.CatResponse{Cat: cat}
	render.JSON(w, r, res)

	catEntry := dbmodel.CatEntry{CatDesc: cat}
	database.DB.Create(&catEntry)

}

