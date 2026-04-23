package main

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		json.NewEncoder(w).Encode([]map[string]interface{}{
			{"id": 1, "name": "Produkt 1"},
			{"id": 2, "name": "Produkt 2"},
			{"id": 3, "name": "Produkt 3"},
		})
	})

	log.Println("server started on :8080")
	http.ListenAndServe(":8080", nil)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
