package Models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	//add other relevant details...
	CustomerID uint `json:"customer_id"`
	RetailerID uint `json:"retailer_id"`
	ProductID uint `json:"product_id"`
	Status bool //true for active
}
func (b *Order) TableName() string {
	return "orders"
}
