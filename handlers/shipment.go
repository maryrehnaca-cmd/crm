package handlers

import (
	"crm/db"
	"html/template"
	"net/http"
)

type Shipment struct {
	ShipmentID   int
	TrackingNo   string
	CustomerID   int
	Source       string
	Destination  string
	Status       string
	Weight       float64
	ShipmentDate string
}

// Display all shipments
/* func ShipmentPage(w http.ResponseWriter, r *http.Request) {

	// Your existing code here

} */

func ShipmentPage(w http.ResponseWriter, r *http.Request) {

	rows, err := db.DB.Query(`
		SELECT shipment_id, tracking_no, customer_id,
		       source, destination, status, weight, shipment_date
		FROM shipments
	`)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var shipments []Shipment

	for rows.Next() {

		var s Shipment

		err := rows.Scan(
			&s.ShipmentID,
			&s.TrackingNo,
			&s.CustomerID,
			&s.Source,
			&s.Destination,
			&s.Status,
			&s.Weight,
			&s.ShipmentDate,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		shipments = append(shipments, s)
	}

	tmpl, err := template.ParseFiles("templates/shipments.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, shipments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Display Add Shipment form
func AddShipmentPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/addShipment.html"))
	tmpl.Execute(w, nil)
}

// Save shipment into database
func AddShipment(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Request", http.StatusMethodNotAllowed)
		return
	}

	tracking := r.FormValue("tracking_no")
	customer := r.FormValue("customer_id")
	source := r.FormValue("source")
	destination := r.FormValue("destination")
	status := r.FormValue("status")
	weight := r.FormValue("weight")
	date := r.FormValue("shipment_date")

	_, err := db.DB.Exec(`
		INSERT INTO shipments
		(tracking_no, customer_id, source, destination, status, weight, shipment_date)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
	`, tracking, customer, source, destination, status, weight, date)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/shipments", http.StatusSeeOther)
}
