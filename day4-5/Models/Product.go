package Models

import (
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
	if product.Quantity>0 {
		product.Status=true
	}
	return nil
}
//GetProductByID ... Fetch only one product by Id
func GetProductByID(product *Product, id string) (err error) {
	if err = Config2.DB.Where("product_id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}
//UpdateProduct ... Update product
func UpdateProduct(product *Product, id string) (err error) {
	var prevProduct Product
	err = DeleteProduct(&prevProduct,id)
	if err!=nil {
		return err
	} else{
		product.ProductID=id
		CreateProduct(product)
	}
	return nil
}

//DeleteProduct ... Delete user
func DeleteProduct(product *Product, id string) (err error) {
	Config2.DB.Where("product_id = ?", id).Delete(product)
	return nil
}

