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
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err.Error())
	} else {
		//creating product in database
		errs:= Models.CreateProduct(product)
		if errs!=nil {
			w.WriteHeader(http.StatusNotFound)
		}
		product.CreatedAt = time.Now().Local()
		productJson,_:=json.Marshal(product)
		_,e:=w.Write(productJson)
		if e!=nil {
			w.WriteHeader(http.StatusNotFound)
		}
		w.Write([]byte(`"message": "product successfully added"`))
		w.WriteHeader(http.StatusOK)
	}
}
func GetOrderHistory(w http.ResponseWriter, r *http.Request){
	var orders []Models.Order
	//get order history
	err := Models.GetAllOrders(&orders)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err.Error())
	} else {
		//output in postman
		ordersJson,_:=json.Marshal(orders)
		_,e:=w.Write(ordersJson)
		if e!=nil {
			w.WriteHeader(http.StatusNotFound)
		}
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
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err.Error())
		return
	} else {
		//output in postman
		productsJson,_:=json.Marshal(product)
		_,e:=w.Write(productsJson)
		if e!=nil{
			w.WriteHeader(http.StatusNotFound)
			return
		}
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
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err.Error())
		return
	} else {
		//updating in database
		perr:= Models.UpdateProduct(product,id)
		if perr !=nil {
			fmt.Println("failed to updated product")
			w.WriteHeader(http.StatusNotFound)
			return
		}
		//display output in postman
		productJson,_:=json.Marshal(product)
		_,e:=w.Write(productJson)
		if e!=nil{
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`"message: product updated"`))
	}
}