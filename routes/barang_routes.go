package routes

import (
	"my-api/controller"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/barang", controller.GetAllBarang).Methods("GET")
    router.HandleFunc("/barang", controller.CreateBarang).Methods("POST")
    router.HandleFunc("/barang/{id}", controller.UpdateBarang).Methods("PUT")
    router.HandleFunc("/barang/{id}", controller.DeleteBarang).Methods("DELETE")

    return router
}
