package Controllers

import (
	"encoding/json"
	"fmt"
	Config2 "freshers-bootcamp/day4-5/Config"
	"freshers-bootcamp/day4-5/Models"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type CustomerActions interface {}

func GetAllProducts(w http.ResponseWriter, r *http.Request){
	var products []Models.Product
	err := Models.GetAllProducts(&products)
	if err != nil {
		panic(err)
		return
	} else {
		w.WriteHeader(http.StatusAccepted)
		errJson,_ := json.Marshal(products)
		_,e:=w.Write(errJson)
		if e!=nil{
			panic(e)
		}
	}
}

func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	//fetching order details
	order :=&Models.Order{}
	err := json.NewDecoder(r.Body).Decode(order)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
		return
	} else {
		var product Models.Product
		perr:= Models.GetProductByID(&product, order.ProductID)
		if perr != nil{
			fmt.Println("product not found")
			panic(perr)
		}else {
			// checking status of customer
			// to subtract from products and adding delay for customer
			if product.Quantity > 0 {
				product.Quantity--
				//updating products status
				if product.Quantity == 0 {
					product.Status = false
				}
			} else {
				w.Write([]byte(`"message: Can't place order"`))
				return
			}
			var customer Models.Customer
			cerr := Models.GetCustomerByID(&customer, string(order.CustomerID))
			if err != nil {
				product.Quantity++
				panic(cerr)
			} else {
				customer.UpdatedAt = time.Now().Local()
			}
		}

		errs:= Models.CreateOrder(order)
		if errs!=nil {
			panic(errs)
		}
		w.WriteHeader(http.StatusAccepted)
		orderJson,_ := json.Marshal(order)
		_,e:=w.Write(orderJson)
		_,e=w.Write([]byte(`"message: orderPlaced"`))
		if e!=nil{
			panic(e)
		}
	}

}

func ViewPurchaseHistory(w http.ResponseWriter, r *http.Request) (err error){
	//taking ID parameter from url
	vars := mux.Vars(r)
	id:= vars["id"]
	if err = Config2.DB.Where("id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderByID(w http.ResponseWriter, r *http.Request){
	//taking ID parameter from url
	vars := mux.Vars(r)
	id:= vars["id"]
	var order Models.Order
	//get orders from database
	err := Models.GetOrderByID(&order, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	} else {
		//output in postman
		orderJson,_:=json.Marshal(order)
		_,e:=w.Write(orderJson)
		if e!=nil{
			panic(e)
		}
		w.WriteHeader(http.StatusOK)
	}
}