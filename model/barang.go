package model

type Barang struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
	Stok int    `json:"stok"`
}
