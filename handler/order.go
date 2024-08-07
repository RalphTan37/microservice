package handler

import (
	"fmt"
	"net/http"
)

// first order type
type Order struct {
}

// creates an order
func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create an order")
}

// returns all of the elements that match a filter
func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("list all orders")
}

// return an order by its ID
func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get an order by id")
}

// updates an order by its ID
func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update an order by ID")
}

// removes orders by ID
func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete an order by ID")
}
