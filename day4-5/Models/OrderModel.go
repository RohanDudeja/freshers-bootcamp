package Models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	//add other relevant details...
	CustomerID uint `gorm:"customer_id"`
	RetailerID uint `gorm:"retailer_id"`
	ProductID uint `gorm:"product_id"`
	Status bool `gorm:"status"` //true for active
}
func (b *Order) TableName() string {
	return "orders"
}
