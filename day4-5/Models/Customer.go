package Models

import (
	"fmt"
	Config2 "freshers-bootcamp/day4-5/Config"
)

//GetAllUsers Fetch all user data
func GetAllCustomers(customer *[]Customer) (err error) {
	if err = Config2.DB.Find(customer).Error; err != nil {
		return err
	}
	return nil
}

//CreateCustomer ... Insert New data
func CreateCustomer(customer *Customer) (err error) {
	if err = Config2.DB.Create(customer).Error; err != nil {
		return err
	}
	return nil
}
//GetProductByID ... Fetch only one product by Id
func GetCustomerByID(customer *Customer, id string) (err error) {
	if err = Config2.DB.Where("id = ?", id).First(customer).Error; err != nil {
		return err
	}
	return nil
}
//UpdateProduct ... Update product
func UpdateCustomer(customer *Customer, id string) (err error) {
	fmt.Println(customer)
	Config2.DB.Save(customer)
	return nil
}

//DeleteProduct ... Delete user
func DeleteCustomer(customer *Customer, id string) (err error) {
	Config2.DB.Where("id = ?", id).Delete(customer)
	return nil
}


