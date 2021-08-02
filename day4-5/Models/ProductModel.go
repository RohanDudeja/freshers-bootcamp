package Models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	ProductID string `json:"product_id"`
	Name string `json:"Name"`
	RetailerID string `json:"retailer_id"`
	Price uint `json:"price"`
	Quantity uint `json:"quantity"` //quantity available
	Status bool `json:"status"`//true for in stock
}
func (b *Product) TableName() string {
	return "products"
}