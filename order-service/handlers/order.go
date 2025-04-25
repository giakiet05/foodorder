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
	var req OrderRequest
	json.NewDecoder(r.Body).Decode(&req)

	authHeader := r.Header.Get("Authorization")
	userId, err := auth.ParseTokenFromHeader(authHeader)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userExists := client.CheckUserExists(userId)
	foodExists := client.CheckFoodExists(req.FoodId)

	if !userExists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if !foodExists {
		http.Error(w, "Food not found", http.StatusNotFound)
		return
	}

	order := models.Order{
		UserId: userId,
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
