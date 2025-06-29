package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
    dsn := "root:eEfeVkAEAGxKMrOkcAvCatnSBUPasTVx@tcp(interchange.proxy.rlwy.net:12313)/railway"

    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("Gagal koneksi database:", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("Tidak bisa ping database:", err)
    }

    log.Println("Koneksi ke database Railway berhasil!")
}
