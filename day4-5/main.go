package main

import (
	"fmt"
	Config2 "freshers-bootcamp/day4-5/Config"
	Controllers2 "freshers-bootcamp/day4-5/Controllers"
	Models2 "freshers-bootcamp/day4-5/Models"
	mux2 "github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

var err error
func main() {
	Config2.DB, err = gorm.Open("mysql", Config2.DbURL(Config2.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config2.DB.Close()
	Config2.DB.AutoMigrate(&Models2.Product{})
	Config2.DB.AutoMigrate(&Models2.Customer{})
	Config2.DB.AutoMigrate(&Models2.Order{})

	r := mux2.NewRouter()

	api := r.PathPrefix("").Subrouter()
	api.HandleFunc("/product",Controllers2.AddProduct).Methods(http.MethodPost)
	api.HandleFunc("/product/{id}",Controllers2.UpdateProduct).Methods(http.MethodPatch)
	api.HandleFunc("/product/{id}",Controllers2.GetProductByID).Methods(http.MethodGet)
	api.HandleFunc("/products", Controllers2.GetAllProducts).Methods(http.MethodGet)
	api.HandleFunc("/order",Controllers2.PlaceOrder).Methods(http.MethodPost)
	api.HandleFunc("/order/{id}",Controllers2.GetOrderByID).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}