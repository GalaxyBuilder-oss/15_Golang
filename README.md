# README

## About

This is the official Wails React-TS template.

You can configure the project by editing `wails.json`. More information about the project settings can be found
here: https://wails.io/docs/reference/project-config

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.

Untuk membuat REST API menggunakan Go (atau disebut juga Golang), kita bisa menggunakan beberapa framework atau library yang populer seperti Gin, Echo, atau net/http bawaan dari Go sendiri. Di bawah ini saya akan tunjukkan cara membuat REST API sederhana menggunakan net/http, yang merupakan cara paling dasar untuk membuat API di Go.

### Langkah-langkah untuk Membuat REST API dengan Go

1. **Instalasi Go**

Pastikan Anda telah menginstal Go di komputer Anda. Anda bisa mengunduh installer dari [Golang official website](https://golang.org/) dan ikuti petunjuk instalasinya.

2. **Buat Struktur Proyek**

Buat sebuah direktori untuk proyek Anda dan buat file utama `main.go`.

3. **Buat API dengan net/http**

Berikut adalah contoh sederhana membuat REST API dengan menggunakan net/http di Go:

```go
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
```

4. **Menjalankan Aplikasi**

Simpan file `main.go`, dan dari terminal, masuk ke direktori proyek dan jalankan dengan perintah:

```sh
go run main.go
```

API ini akan berjalan di `http://localhost:8080`. Anda dapat mengakses endpoint:
- `GET /users` untuk mendapatkan semua pengguna
- `GET /users/{id}` untuk mendapatkan pengguna berdasarkan ID

### Penjelasan Kode

- **`User` struct**: Digunakan untuk mendefinisikan struktur data pengguna dengan field ID, Name, dan Email.
- **`users` variable**: Contoh data pengguna yang akan digunakan untuk API.
- **Routing**: Menggunakan `http.HandleFunc` untuk menentukan handler untuk setiap endpoint.
- **Handler functions**:
  - `getUsers`: Handler untuk endpoint `GET /users` yang mengembalikan semua pengguna dalam format JSON.
  - `getUserByID`: Handler untuk endpoint `GET /users/{id}` yang mengembalikan pengguna berdasarkan ID yang diberikan dalam URL.

Dengan menggunakan net/http dan mengikuti struktur di atas, Anda dapat membuat REST API yang sederhana dan efisien di Go. Framework seperti Gin atau Echo juga bisa dipertimbangkan untuk keperluan yang lebih kompleks atau meningkatkan kecepatan pengembangan.