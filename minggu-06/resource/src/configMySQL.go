package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Konfigurasi koneksi MySQL
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/database_name")
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
	rows, err := db.Query("SELECT * FROM nama_tabel")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Membaca hasil query
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(id, name)
	}
}
