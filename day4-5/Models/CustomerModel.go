package Models

import "github.com/jinzhu/gorm"

type Customer struct {
	gorm.Model
	Name string `json:"Name"`
	Address string `json:"Address"`
}
func (b *Customer) TableName() string {
	return "customers"
}