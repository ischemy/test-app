package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "://github.com"
)

func main() {
	dbHost := os.Getenv("DB_HOST") // 10.192.15.6
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbUser, dbPass, dbName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := db.Ping()
		if err != nil {
			fmt.Fprintf(w, "Status: Gagal koneksi ke Lintasdbapp1! Error: %v", err)
			return
		}
		fmt.Fprint(w, "Status: Berhasil Terhubung ke PostgreSQL di 10.192.15.6")
	})

	log.Println("Aplikasi berjalan di port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
