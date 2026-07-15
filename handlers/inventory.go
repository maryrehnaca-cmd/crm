package handlers

import (
	"crm/db"
	"html/template"
	"net/http"
)

type Inventory struct {
	ProductID   int
	ProductName string
	Quantity    int
	Warehouse   string
	UnitPrice   float64
}

// Display Inventory
func InventoryPage(w http.ResponseWriter, r *http.Request) {

	rows, err := db.DB.Query(`
		SELECT product_id,
       product_name,
       quantity,
       warehouse,
       unit_price
FROM inventory
ORDER BY product_id
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []Inventory

	for rows.Next() {

		var i Inventory

		err := rows.Scan(
			&i.ProductID,
			&i.ProductName,
			&i.Quantity,
			&i.Warehouse,
			&i.UnitPrice,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		items = append(items, i)
	}

	tmpl := template.Must(template.ParseFiles("templates/inventory.html"))
	tmpl.Execute(w, items)
}

// Display Add Inventory Page
func AddInventoryPage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/addInventory.html"))
	tmpl.Execute(w, nil)

}

// Save Inventory
func AddInventory(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Request", http.StatusMethodNotAllowed)
		return
	}

	productName := r.FormValue("product_name")
	quantity := r.FormValue("quantity")
	warehouse := r.FormValue("warehouse")
	unitPrice := r.FormValue("unit_price")

	_, err := db.DB.Exec(`
    INSERT INTO inventory
    (product_name, quantity, warehouse, unit_price)
    VALUES ($1,$2,$3,$4)
`, productName, quantity, warehouse, unitPrice)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/inventory", http.StatusSeeOther)
}
