package controllers

import (
	"encoding/json"
	"mvc-gorm/models"
	"mvc-gorm/repositories"
	"net/http"
	"strconv"
	"time"
)

func GetCategory(w http.ResponseWriter, r *http.Request){
	categories, err := repositories.GetCategory()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(categories)
}

func GetCategoryByID(w http.ResponseWriter, r *http.Request){
	idParams := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idParams)
	if err != nil {
		http.Error(w, "product not found", http.StatusBadRequest)
		return
	}

	category, err := repositories.GetCategoryByID(id)
	if err !=nil {
		http.Error(w, "product not found", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(category)
}

func CreateCategory(w http.ResponseWriter, r *http.Request){
	var category models.Category
	json.NewDecoder(r.Body).Decode(&category)

	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()

	err:= repositories.CreateCategory(category)
	if err !=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)
}



