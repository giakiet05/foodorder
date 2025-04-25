package handlers

import (
	"encoding/json"
	"github.com/giakiet05/foodorder/food-service/db"
	"github.com/giakiet05/foodorder/food-service/models"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateFood(w http.ResponseWriter, r *http.Request) {
	var food models.Food

	json.NewDecoder(r.Body).Decode(&food)

	db.DB.Create(&food)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(food)
}

func GetFoods(w http.ResponseWriter, r *http.Request) {
	var foods []models.Food
	db.DB.Find(&foods)
	json.NewEncoder(w).Encode(&foods)
}

func GetFood(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var food models.Food
	if err := db.DB.First(&food, id).Error; err != nil {
		http.Error(w, "Food not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(food)
}
