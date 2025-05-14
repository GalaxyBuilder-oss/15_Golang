package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// User struct untuk menyimpan informasi pengguna
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Data pengguna contoh
var users = []User{
	{ID: 1, Name: "John Doe", Email: "john@example.com"},
	{ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
}

func main() {
	// Routing
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/users/", getUserByID)

	// Menjalankan server
	log.Fatal(http.ListenAndServe(":8080", nil))
	
}

// Handler untuk endpoint GET /users
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Handler untuk endpoint GET /users/{id}
func getUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Mendapatkan parameter ID dari URL
	id := r.URL.Path[len("/users/"):]

	// Mencari pengguna berdasarkan ID
	for _, user := range users {
		if id == string(user.ID) {
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	// Mengirim respons jika pengguna tidak ditemukan
	http.NotFound(w, r)
}
