package Models

import "github.com/jinzhu/gorm"

type Customer struct {
	gorm.Model
	CustomerID string `json:"customer_id"`
	Name string `json:"Name"`
	Address string `json:"Address"`
}
func (b *Customer) TableName() string {
	return "customers"
}