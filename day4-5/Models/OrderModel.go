package Models

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	//add other relevant details...
	OrderID string `json:"order_id"`
	CustomerID string `json:"customer_id"`
	RetailerID string `json:"retailer_id"`
	ProductID string `json:"product_id"`
	Status bool //true for active
}
func (b *Order) TableName() string {
	return "orders"
}
