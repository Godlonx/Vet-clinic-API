package cat

import (
	"clinic/config"
	"clinic/database/dbmodel"
	"clinic/pkg/helper"
	model "clinic/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type CatConfig struct {
	*config.Config
}

func New(configuration *config.Config) *CatConfig {
	return &CatConfig{configuration}
}

func (config *CatConfig) CreateCat(w http.ResponseWriter, r *http.Request) {
	req := &model.CatRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	res := &model.CatResponse{Age: req.Age, Breed: req.Breed, Name: req.Name, Weight: req.Weight}
	render.JSON(w, r, res)

	catEntry := dbmodel.Cat{Age: req.Age, Breed: req.Breed, Name: req.Name, Weight: req.Weight}
	config.CatRepository.Create(&catEntry)
}

func (config *CatConfig) GetCats(w http.ResponseWriter, r *http.Request) {
	cats, err := config.CatRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Error fetching cats"})
		return
	}
	render.JSON(w, r, cats)
}

func (config *CatConfig) GetCat(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strId)
	if err != nil || id < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	cat, err := config.CatRepository.Find(id)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Error fetching cat"})
		return
	}
	render.JSON(w, r, cat)
}

func (config *CatConfig) UpdateCat(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	catId, err := strconv.Atoi(strId)
	if err != nil || catId < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	cat, err := config.CatRepository.Find(catId)
	if err != nil {
		http.Error(w, "Failed to retrieve a cat on this id", http.StatusInternalServerError)
		return
	}

	var data map[string]interface{}

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Cannot decode body", http.StatusInternalServerError)
		return
	}

	helper.ApplyChanges(data, cat)

	cat, err = config.CatRepository.Update(cat)
	if err != nil {
		http.Error(w, "Failed to update cat on this id", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, cat)
}

func (config *CatConfig) DeleteCat(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	catId, err := strconv.Atoi(strId)
	if err != nil || catId < 0 {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	err = config.CatRepository.Delete(catId)
	if err != nil {
		http.Error(w, "Failed to retrieve visits for this catId", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]string{"message": "Cat deleted"})
}
