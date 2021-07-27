package Models

import (
	"fmt"
	Config2 "freshers-bootcamp/day4-5/Config"
)

//GetAllUsers Fetch all order data
func GetAllOrders(order *[]Order) (err error) {
	if err = Config2.DB.Find(order).Error; err != nil {
		return err
	}
	return nil
}

//CreateOrder ... Place New order
func CreateOrder(order *Order) (err error) {
	if err = Config2.DB.Create(order).Error; err != nil {
		return err
	}
	return nil
}
//GetProductByID ... Fetch only one order by Id
func GetOrderByID(order *Order, id string) (err error) {
	if err = Config2.DB.Where("id = ?", id).First(order).Error; err != nil {
		return err
	}
	return nil
}
//UpdateOrder ... from placed to delievered

func UpdateOrder(order *Order, id string) (err error) {
	fmt.Println(order)
	Config2.DB.Save(order)
	return nil
}

