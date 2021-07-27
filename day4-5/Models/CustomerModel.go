package Models

import "github.com/jinzhu/gorm"

type Customer struct {
	gorm.Model
	Name string `gorm:"Name"`
	Address string `gorm:"Address"`
}
func (b *Customer) TableName() string {
	return "customers"
}