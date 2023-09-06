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
	// llamar a nuestras rutas
	mux.HandleFunc(prefijo+"ejemplo", handlers.Ejemplo_get).Methods("GET")
	mux.HandleFunc(prefijo+"ejemplo/{id:[0-9]+", handlers.Ejemplo_get_con_parametros).Methods("GET")
	mux.HandleFunc(prefijo+"ejemplo", handlers.Ejemplo_post).Methods("POST")
	mux.HandleFunc(prefijo+"ejemplo/{id:[0-9]+", handlers.Ejemplo_put).Methods("PUT")
	mux.HandleFunc(prefijo+"ejemplo/{id:[0-9]+", handlers.Ejemplo_delete).Methods("DELETE")
	mux.HandleFunc(prefijo+"query-string", handlers.Ejemplo_get_querystring).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", mux))
}