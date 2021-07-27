package Models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name string `json:"Name"`
	RetailerID uint `json:"retailer_id"`
	Price uint `json:"price"`
	Quantity uint `json:"quantity"` //quantity available
	Status bool //true for in stock
}
func (b *Product) TableName() string {
	return "products"
}