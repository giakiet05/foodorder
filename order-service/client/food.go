package client

import (
	"fmt"
	"net/http"
)

func CheckFoodExists(foodId uint) bool {
	res, err := http.Get(fmt.Sprintf("http://localhost:8001/foods/%d", foodId))
	if err != nil || res.StatusCode != 200 {
		return false
	}
	return true
}
