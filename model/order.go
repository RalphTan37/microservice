package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct { //order properties
	//json tags - enable to encode + decode to json using the standard library
	OrderID     uint64     `json:"order_id"`     //represent the unique identifier of the order
	CustomerID  uuid.UUID  `json:"customer_id"`  //universally unique identifier
	LineItems   []LineItem `json:"line_items"`   //line of items
	CreatedAt   *time.Time `json:"created_at"`   //create status
	ShippedAt   *time.Time `json:"shipped_at"`   //ship status
	CompletedAt *time.Time `json:"completed_at"` //complete status
}

type LineItem struct { //represent the individual item purchases w/in an order
	ItemID   uuid.UUID `json:"item_id"`  //represents which item it is
	Quantity uint      `json:"quantity"` //represents how many of these items were purchased
	Price    uint      `json:"price"`    //represents the price of a single unit
}
