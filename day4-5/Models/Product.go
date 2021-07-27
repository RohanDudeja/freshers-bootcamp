package Models

import (
"fmt"
Config2 "freshers-bootcamp/day4-5/Config"
_ "github.com/go-sql-driver/mysql"
)
//GetAllUsers Fetch all user data
func GetAllProducts(product *[]Product) (err error) {
	if err = Config2.DB.Find(product).Error; err != nil {
		return err
	}
	return nil
}

//CreateProduct ... Insert New data
func CreateProduct(product *Product) (err error) {
	if err = Config2.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}
//GetProductByID ... Fetch only one product by Id
func GetProductByID(product *Product, id string) (err error) {
	if err = Config2.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}
//UpdateProduct ... Update product
func UpdateProduct(product *Product, id string) (err error) {
	fmt.Println(product)
	Config2.DB.Save(product)
	return nil
}

//DeleteProduct ... Delete user
func DeleteProduct(product *Product, id string) (err error) {
	Config2.DB.Where("id = ?", id).Delete(product)
	return nil
}

