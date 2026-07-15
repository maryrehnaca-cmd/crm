package main

import (
	"crm/db"
	"crm/handlers"
	"log"
	"net/http"
)

func main() {

	db.ConnectDB()

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.LoginPage)
	http.HandleFunc("/signup", handlers.SignupPage)
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/shipments", handlers.ShipmentPage)
	//http.HandleFunc("/shipments", handlers.ShipmentPage)
	http.HandleFunc("/addShipmentPage", handlers.AddShipmentPage)
	http.HandleFunc("/addShipment", handlers.AddShipment)
	http.HandleFunc("/customers", handlers.CustomerPage)
	http.HandleFunc("/addCustomerPage", handlers.AddCustomerPage)
	http.HandleFunc("/addCustomer", handlers.AddCustomer)
	http.HandleFunc("/inventory", handlers.InventoryPage)
	http.HandleFunc("/addInventoryPage", handlers.AddInventoryPage)
	http.HandleFunc("/addInventory", handlers.AddInventory)

	log.Println("Server Running :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
