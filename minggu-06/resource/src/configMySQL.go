package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Konfigurasi koneksi MySQL
	username := "root"
	password := "admin"
	host := "localhost"
	port := "3306" // Ganti dengan port yang sesuai
	databaseName := "db_mahasiswa"

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, databaseName)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Membuka koneksi database
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	// Query untuk membaca data dari tabel MySQL
	rows, err := db.Query("SELECT id, nama, nim, email, kelas FROM mahasiswati")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Membaca hasil query
	for rows.Next() {
		var id, nim, kelas int
		var nama, email string
		err = rows.Scan(&id, &nama, &nim, &email, &kelas)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("ID: %d, Nama: %s, NIM: %d, Email: %s, Kelas: %s\n", id, nama, nim, email, kelas)
	}
}
