package handlers

import (
	"crm/db"
	"html/template"
	"net/http"
)

type Customer struct {
	CustomerID  int
	Name        string
	CompanyName string
	Phone       string
	Email       string
	Address     string
}

func CustomerPage(w http.ResponseWriter, r *http.Request) {

	rows, err := db.DB.Query(`
	SELECT customer_id,
	       name,
	       company_name,
	       phone,
	       email,
	       address
	FROM customers
	ORDER BY customer_id
`)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var customers []Customer

	for rows.Next() {
		var c Customer

		err := rows.Scan(
			&c.CustomerID,
			&c.Name,
			&c.CompanyName,
			&c.Phone,
			&c.Email,
			&c.Address,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		customers = append(customers, c)
	}

	tmpl := template.Must(template.ParseFiles("templates/customers.html"))
	tmpl.Execute(w, customers)
}

// Show Add Customer Form
func AddCustomerPage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/addCustomer.html"))
	tmpl.Execute(w, nil)

}

// Save Customer
func AddCustomer(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Request", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	company := r.FormValue("company_name")
	phone := r.FormValue("phone")
	email := r.FormValue("email")
	address := r.FormValue("address")

	_, err := db.DB.Exec(`
		INSERT INTO customers
(name, company_name, phone, email, address)
VALUES ($1,$2,$3,$4,$5)
	`, name, company, phone, email, address)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/customers", http.StatusSeeOther)

}
