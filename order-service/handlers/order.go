package handlers

import (
	"encoding/json"
	"github.com/giakiet05/foodorder/order-service/auth"
	"github.com/giakiet05/foodorder/order-service/client"
	"github.com/giakiet05/foodorder/order-service/db"
	"github.com/giakiet05/foodorder/order-service/models"
	"github.com/gorilla/mux"
	"net/http"
)

type OrderRequest struct {
	FoodId uint `json:"food_id"`
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	tokenStr := r.Header.Get("Authorization")
	if tokenStr == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}

	claims, err := auth.ParseToken(tokenStr)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	var req OrderRequest
	json.NewDecoder(r.Body).Decode(&req)

	if !client.CheckFoodExists(req.FoodId) {
		http.Error(w, "Food does not exist", http.StatusBadRequest)
		return
	}

	order := models.Order{
		UserId: claims.UserId,
		FoodId: req.FoodId,
		Status: "pending",
	}
	db.DB.Create(&order)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	var orders []models.Order
	db.DB.Find(&orders)
	json.NewEncoder(w).Encode(orders)

}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var order models.Order

	if err := db.DB.Find(&order, id).Error; err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(order)
}
