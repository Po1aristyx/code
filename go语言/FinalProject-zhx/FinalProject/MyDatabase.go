package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

var db *sql.DB

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type City struct {
	ID          int    `json: "id"`
	Name        string `json: "name"`
	CountryCode string `json: "CountryCode"`
	District    string `json: "District"`
	Population  int    `json: "Population"`
}

func initDb() {

}

func getProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Not found", http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getCitys(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Not found", http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("select ID,Name,CountryCode,District, Population from city  where ID=(select max(ID) from city)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var citys []City
	for rows.Next() {
		var c City
		if err := rows.Scan(&c.ID, &c.Name, &c.CountryCode, &c.District, &c.Population); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		citys = append(citys, c)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(citys); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
