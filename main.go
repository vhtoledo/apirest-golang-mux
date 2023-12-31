package main

import (
	"golang-mux-apirest/handlers"
	"golang-mux-apirest/middleware"
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

	mux.HandleFunc(prefijo+"productos", handlers.Productos_get).Methods("GET")
	mux.HandleFunc(prefijo+"productos/{id:[0-9]+}", handlers.Productos_get_con_parametro).Methods("GET")
	mux.HandleFunc(prefijo+"productos", handlers.Productos_post).Methods("POST")
	mux.HandleFunc(prefijo+"productos/{id:[0-9]+}", handlers.Productos_put).Methods("PUT")
	mux.HandleFunc(prefijo+"productos/{id:[0-9]+}", handlers.Productos_delete).Methods("DELETE")

	mux.HandleFunc(prefijo+"productos-fotos/{id:[0-9]+}", handlers.ProductosFotosUpload).Methods("POST")
	mux.HandleFunc(prefijo+"productos-fotos/{id:[0-9]+}", handlers.ProductosFotos_get_por_producto).Methods("GET")
	mux.HandleFunc(prefijo+"productos-fotos/{id:[0-9]+}", handlers.ProductosFotosDelete).Methods("DELETE")

	mux.HandleFunc(prefijo+"seguridad/registro", handlers.Seguridad_registro).Methods("POST")
	mux.HandleFunc(prefijo+"seguridad/login", handlers.Seguridad_login).Methods("POST")
	mux.HandleFunc(prefijo+"seguridad/protegido", middleware.ValidarJWT(handlers.Seguridad_protegido)).Methods("GET")

	//cors
	handler := cors.AllowAll().Handler(mux)
	//log.Fatal(http.ListenAndServe(":8080", mux))
	log.Fatal(http.ListenAndServe(":8080", handler))
}