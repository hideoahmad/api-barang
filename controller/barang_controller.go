package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"my-api/config"
	"my-api/model"

	"github.com/gorilla/mux"
)

// GET /barang
func GetAllBarang(w http.ResponseWriter, r *http.Request) {
    rows, err := config.DB.Query("SELECT id, nama, stok FROM barang")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var barangs []model.Barang
    for rows.Next() {
        var b model.Barang
        if err := rows.Scan(&b.ID, &b.Nama, &b.Stok); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        barangs = append(barangs, b)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(barangs)
}

// POST /barang
func CreateBarang(w http.ResponseWriter, r *http.Request) {
    var b model.Barang
    err := json.NewDecoder(r.Body).Decode(&b)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    res, err := config.DB.Exec("INSERT INTO barang (nama, stok) VALUES (?, ?)", b.Nama, b.Stok)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    id, _ := res.LastInsertId()
    b.ID = int(id)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(b)
}

// PUT /barang/{id}
func UpdateBarang(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    var b model.Barang
    err := json.NewDecoder(r.Body).Decode(&b)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err = config.DB.Exec("UPDATE barang SET nama=?, stok=? WHERE id=?", b.Nama, b.Stok, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    b.ID = id
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(b)
}

// DELETE /barang/{id}
func DeleteBarang(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    _, err := config.DB.Exec("DELETE FROM barang WHERE id=?", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Barang dengan ID %d berhasil dihapus", id)
}
