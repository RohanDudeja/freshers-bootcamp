# Hypothetical E-Commerece Server API
## Retailer Service Requirements

Build a service to support a hypothetical retailer with the following requirements:

1. The retailer has some products which he wants to sell. The retailer should be able to add a product, update the available quantity and price of the product as and when required.
2. Customers should be able to view all the available products along with their prices using the REST API. 
3. A customer should be able to place an order. And later should be able to view his order history.
4. The retailer should be able to view a detailed history of all the transactions of his business.
5. Since our retailer has many customers, He wants you to design the solution in such a way that once an order of the customer is executed, there should be a gap or cool-down period of 5 mins for the next order to be executed. (Cool down is at customer level)

Tech-Stack used:
* mysql, gorm
* gorilla-mux, gin

## Api contract:

1. POST /product  
request:  
```json
{
	"product_name": "bottle",
	"price": 50,
	"quantity": 40
}
```
response:  
```json
{
	"id": "PROD12345",
	"product_name": "bottle",
	"price": 50,
	"quantity": 40,
	"message": "product successfully added"
}
```
2. PATCH /product/:id  
request:  
```json
{
	"price": 60,
	"quantity": 4
}
```
response:  
```json
{
	"id": "PROD12345",
	"product_name": "bottle",
	"price": 60,
	"quantity": 4
}
```
3. GET /product/:id  
response:  
```json
{
	"id": "PROD12345",
	"product_name": "bottle",
	"price": 60,
	"quantity": 4
}
```
4. GET /products  
response:
```json
{
	"products": [{
			"id": "PROD12345",
			"product_name": "bottle",
			"price": 60,
			"quantity": 4
		},
		{
			"id": "PROD45545",
			"product_name": "basket",
			"price": 100,
			"quantity": 4
		}
	]
}
```
5. POST /order  
request:  
```json
{
      "customer_id": "CST12345",
	"product_id": "PROD12345",
	"quantity": 1
}
```
response:  
```json
{
	"id": "ORD1234",
	"product_id": "PROD12345",
	"quantity": 1,
	"status": "order placed" # status can be "order placed" or "processed" or "failed"
}
```

6. GET /order/:id  
response:  
```json
{
	"id": "ORD1234",
	"product_id": "PROD12345",
	"quantity": 1,
	"status": "order placed" # status can be "order placed" or "processed" or "failed"
}
```



<img width="903" alt="Screenshot 2021-07-27 at 5 10 25 PM" src="https://user-images.githubusercontent.com/43816495/127148082-abb3a306-45a1-4efa-b762-1e1ef5be596b.png">

Tasks available
*Retailer:*
> 1. AddProduct
> 2. UpdateProduct
> 3. CheckOrderHistory [depending on Retailer_id]

*Customer:*
> 1. PlaceOrder
> 2. CheckOrderHistory [depending on Customer_id]

I tried to implement on basis of MVC framework:
*Models and Tables are defined for:*
> Customers, 
> Products,
> Orders

*Controllers are defined for:*
> Retailer,
> Customer

Command to run the code:
1. Clone the repo
2. Change the database and password for mysql in `freshers-bootcamp/day4_5/Config/Database.go`
3. and run `go run freshers-bootcamp/day4-5/main.go`

*API calls*
* POST: `retailer/product` to AddProduct
* PATCH: `reatiler/product/{id}` to UpdateProduct
* GET: `/consumer/product/{id}` to GetProductByID
* GET:	`consumer/products` to GetAllProducts
* POST:	`consumer/order` to PlaceOrder
* GET: `consumer/order/{id}` to GetOrderByID




