package cat

import (
	"net/http"
    "github.com/go-chi/render"
	"clinic/config"
	"clinic/pkg/models"
	"clinic/database/dbmodel"
)

type CatConfig struct {
    *config.Config
}

func New(configuration *config.Config) *CatConfig {
    return &CatConfig{configuration}
}

func (config *CatConfig) CatHandler(w http.ResponseWriter, r *http.Request){
	req := &model.CatRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error":"Invalid request payload"})
		return
	}


	res := &model.CatResponse{Age: req.Age, Breed: req.Breed, Name: req.Name, Weight: req.Weight}
	render.JSON(w, r, res)

	catEntry := dbmodel.Cat{ Age: req.Age, Breed: req.Breed, Name: req.Name, Weight: req.Weight}
    config.CatRepository.Create(&catEntry)

	
}

