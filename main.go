package main

import (
	"golang-mux-apirest/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main () {
	mux := mux.NewRouter()
	prefijo := "/api/v1"
	mux.HandleFunc(prefijo+"ejemplo", handlers.Ejemplo_get).Methods("GET")
	//mux.HandleFunc(prefijo+"ejemplo", handlers.Ejemplo_get).Methods("GET")
	mux.HandleFunc(prefijo+"ejemplo", handlers.Ejemplo_post).Methods("POST")
	mux.HandleFunc(prefijo+"ejemplo", handlers.Ejemplo_put).Methods("PUT")
	mux.HandleFunc(prefijo+"ejemplo", handlers.Ejemplo_delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", mux))
}