# Following is the server implementation for hypothetical retailer
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



