package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
    dbUser := os.Getenv("MYSQLUSER")
    dbPass := os.Getenv("MYSQLPASSWORD")
    dbHost := os.Getenv("MYSQLHOST")
    dbPort := os.Getenv("MYSQLPORT")
    dbName := os.Getenv("MYSQLDATABASE")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

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
