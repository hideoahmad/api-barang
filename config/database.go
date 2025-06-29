package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
    var err error
    DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_barang")
    if err != nil {
        log.Fatal("Gagal koneksi DB:", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("DB tidak bisa diakses:", err)
    }

    log.Println("Koneksi DB sukses!")
}
