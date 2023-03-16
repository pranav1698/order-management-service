package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	Orders []Order
)

type Item struct {
	Id          string  `json:"id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

type Order struct {
	Id           string  `json:"id"`
	Status       string  `json:"status"`
	Items        []Item  `json:"items"`
	Total        float64 `json:"total"`
	CurrencyUnit string  `json:"currencyUnit"`
}

// GetOrders fetches all Orders present in memory
// Example: http://localhost:8080/Orders
func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(Orders) != 0 {
		json.NewEncoder(w).Encode(Orders)
	}

}

// CreateNewOrder creates a new order coming in the JSON request and stores it in memory
// Example: http://localhost:8080/create
func CreateNewOrder(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Cannot create order, due to error: %s", err)
		return
	}

	var order Order
	json.Unmarshal(reqBody, &order)
	log.Println(order)

	if !CheckDuplicateOrder(order) {
		Orders = append(Orders, order)
		json.NewEncoder(w).Encode(order)

		newData, err := json.Marshal(Orders)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(newData))
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Duplicate Order, order already present in storage")
		return
	}

}

// Update Status will update the order given the id and status of the order coming in the JSON request
// Example: http://localhost:8080/update
func UpdateStatus(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var updatedOrder Order
	json.Unmarshal(reqBody, &updatedOrder)

	for index, order := range Orders {
		if updatedOrder.Id == order.Id {
			Orders = append(Orders[:index], Orders[index+1:]...)

			order.Status = updatedOrder.Status
			Orders = append(Orders, order)

			json.NewEncoder(w).Encode(Orders)
			return
		}
	}

	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintln(w, "Order ID does not exits")
}

// Get Order filters the Orders according to filter mentioned in the URL
// Example: http://localhost:8080/Orders/id/abcdef-123456
func GetOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	filter := params["filter"]
	value := params["value"]
	var filteredOrders []Order

	switch filter {
	case "id":
		for _, order := range Orders {
			if order.Id == value {
				filteredOrders = append(filteredOrders, order)
			}
		}
		json.NewEncoder(w).Encode(filteredOrders)
	case "status":
		for _, order := range Orders {
			if order.Status == value {
				filteredOrders = append(filteredOrders, order)
			}
		}
		json.NewEncoder(w).Encode(filteredOrders)
	case "total":
		total, _ := strconv.ParseFloat(value, 64)
		for _, order := range Orders {

			if order.Total == total {
				filteredOrders = append(filteredOrders, order)
			}
		}
		json.NewEncoder(w).Encode(filteredOrders)
	case "currencyUnit":
		for _, order := range Orders {
			if order.CurrencyUnit == value {
				filteredOrders = append(filteredOrders, order)
			}
		}
		json.NewEncoder(w).Encode(filteredOrders)
	}
}

// checDuplicateOrder will check if there is no duplicate order when we are creating a new order
func CheckDuplicateOrder(newOrder Order) bool {
	for _, order := range Orders {
		if order.Id == newOrder.Id {
			return true
		}
	}

	return false
}
