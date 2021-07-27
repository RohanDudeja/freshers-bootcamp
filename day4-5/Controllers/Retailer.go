package Controllers

import (
	"encoding/json"
	"fmt"
	"freshers-bootcamp/day4-5/Models"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Retailer interface {}

func AddProduct(w http.ResponseWriter, r *http.Request){
	product:= &Models.Product{}
	//decoding input
	err := json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		panic(err)
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err.Error())
	} else {
		//creating product in database
		errs:= Models.CreateProduct(product)
		if err!=nil {
			panic(errs)
		}
		product.CreatedAt = time.Now().Local()
		productJson,_:=json.Marshal(product)
		w.Write(productJson)
		w.WriteHeader(http.StatusOK)
	}
}
func GetOrderHistory(w http.ResponseWriter, r *http.Request){
	var orders []Models.Order
	//get order history
	err := Models.GetAllOrders(&orders)
	if err != nil {
		panic(err)
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err.Error())
	} else {
		//output in postman
		ordersJson,_:=json.Marshal(orders)
		w.Write(ordersJson)
		w.WriteHeader(http.StatusOK)
	}
}
func GetProductByID(w http.ResponseWriter, r *http.Request){
	//taking ID parameter from url
	vars := mux.Vars(r)
	id:= vars["id"]
	var product Models.Product
	// get product list
	err := Models.GetProductByID(&product, id)
	if err != nil {
		panic(err)
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err.Error())
	} else {
		//output in postman
		productsJson,_:=json.Marshal(product)
		w.Write(productsJson)
		w.WriteHeader(http.StatusOK)
	}
}
func UpdateProduct(w http.ResponseWriter, r *http.Request){
	product:= &Models.Product{}
	//taking ID parameter from url
	vars := mux.Vars(r)
	id:= vars["id"]
	err := json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		panic(err)
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err.Error())
	} else {
		product.CreatedAt = time.Now().Local()
		//updating in database
		Models.UpdateProduct(product,id)
		//display output in postman
		productJson,_:=json.Marshal(product)
		w.Write(productJson)
		w.WriteHeader(http.StatusOK)
	}
}