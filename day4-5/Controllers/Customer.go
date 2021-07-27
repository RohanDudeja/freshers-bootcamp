package Controllers

import (
	"encoding/json"
	Config2 "freshers-bootcamp/day4-5/Config"
	"freshers-bootcamp/day4-5/Models"
	"github.com/gorilla/mux"
	"net/http"
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
		w.Write(errJson)
	}
}

func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	order :=&Models.Order{}
	err := json.NewDecoder(r.Body).Decode(order)
	if err != nil {
		panic(err)
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		errs:= Models.CreateOrder(order)
		if errs!=nil {
			panic(errs)
		}
		w.WriteHeader(http.StatusAccepted)
		orderJson,_ := json.Marshal(order)
		w.Write(orderJson)
		w.Write([]byte(`{"message: orderPlaced"}`))
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
		panic(err)
		w.WriteHeader(http.StatusNotFound)
	} else {
		//output in postman
		orderJson,_:=json.Marshal(order)
		w.Write(orderJson)
		w.WriteHeader(http.StatusOK)
	}
}