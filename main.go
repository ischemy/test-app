package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// Konfigurasi diambil dari Docker Compose environment
	dbHost := os.Getenv("DB_HOST")     // 10.192.15.6
	dbUser := os.Getenv("DB_USER")     // postgres
	dbPass := os.Getenv("DB_PASSWORD") 
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbUser, dbPass, dbName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error open connection: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := db.Ping()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "❌ test-app gagal konek ke DB di 10.192.15.6: %v", err)
			return
		}
		fmt.Fprint(w, "✅ test-app berhasil koneksi ke PostgreSQL Lintasdbapp1!")
	})

	log.Println("test-app start on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
