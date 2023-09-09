package main

import (
	"golang-mux-apirest/handlers"
	//"golang-mux-apirest/modelos"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main () {
	//mirar la bd
	//modelos.Migraciones()

	mux := mux.NewRouter()
	prefijo := "/api/v1"
	// rutas
	mux.HandleFunc(prefijo+"ejemplo", handlers.Ejemplo_get).Methods("GET")
	mux.HandleFunc(prefijo+"ejemplo/{id:[0-9]+", handlers.Ejemplo_get_con_parametros).Methods("GET")
	mux.HandleFunc(prefijo+"ejemplo", handlers.Ejemplo_post).Methods("POST")
	mux.HandleFunc(prefijo+"ejemplo/{id:[0-9]+", handlers.Ejemplo_put).Methods("PUT")
	mux.HandleFunc(prefijo+"ejemplo/{id:[0-9]+", handlers.Ejemplo_delete).Methods("DELETE")
	mux.HandleFunc(prefijo+"query-string", handlers.Ejemplo_get_querystring).Methods("GET")
	mux.HandleFunc(prefijo+"upload", handlers.Ejemplo_upload).Methods("POST")
	mux.HandleFunc(prefijo+"archivo", handlers.EjemploVerFoto).Methods("GET")

	mux.HandleFunc(prefijo+"categorias", handlers.Categoria_get).Methods("GET")
	mux.HandleFunc(prefijo+"categorias/{id:[0-9]+}", handlers.Categoria_get_con_parametro).Methods("GET")
	mux.HandleFunc(prefijo+"categorias", handlers.Categoria_post).Methods("POST")
	mux.HandleFunc(prefijo+"categorias/{id:[0-9]+}", handlers.Categoria_put).Methods("PUT")
	mux.HandleFunc(prefijo+"categorias/{id:[0-9]+}", handlers.Categoria_delete).Methods("DELETE")

	//cors
	handler := cors.AllowAll().Handler(mux)
	//log.Fatal(http.ListenAndServe(":8080", mux))
	log.Fatal(http.ListenAndServe(":8080", handler))
}