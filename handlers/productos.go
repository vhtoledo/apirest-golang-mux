package handlers

import (
	"golang-mux-apirest/database"
	"golang-mux-apirest/dto"
	"golang-mux-apirest/modelos"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/gosimple/slug"
)

func Productos_get(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	datos := modelos.Productos{}
	database.Database.Order("id desc").Preload("Categoria").Find(&datos)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(datos)
}
func Productos_get_con_parametro(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	//id, _ = strconv.Atoi(vars["id"])
	datos := modelos.Producto{}
	if err := database.Database.Preload("Categoria").First(&datos, vars["id"]); err.Error != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Recurso no disponible",
		}
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(respuesta)
		return
	} else {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(datos)
	}

}
func Productos_post(response http.ResponseWriter, request *http.Request) {
	var producto dto.ProductoDto
	if err := json.NewDecoder(request.Body).Decode(&producto); err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(respuesta)
		return
	}
	datos := modelos.Producto{Nombre: producto.Nombre, Slug: slug.Make(producto.Nombre), Precio: producto.Precio, Stock: producto.Stock, Descripcion: producto.Descripcion, CategoriaID: producto.CategoriaID, Fecha: time.Now()}
	database.Database.Save(&datos)
	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Se creó el registro exitosamente",
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(respuesta)
}
func Productos_put(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])
	var producto dto.ProductoDto

	if err := json.NewDecoder(request.Body).Decode(&producto); err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(respuesta)
		return
	}
	datos := modelos.Producto{}
	if err := database.Database.First(&datos, id); err.Error != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Recurso no disponible",
		}
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(respuesta)
		return
	} else {
		datos.Nombre = producto.Nombre
		datos.Slug = slug.Make(producto.Nombre)
		datos.Precio = producto.Precio
		datos.Stock = producto.Stock
		datos.Descripcion = producto.Descripcion
		datos.CategoriaID = producto.CategoriaID
		database.Database.Save(&datos)
		respuesta := map[string]string{
			"estado":  "ok",
			"mensaje": "Se modificó el registro exitosamente",
		}
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusCreated)
		json.NewEncoder(response).Encode(respuesta)
	}
}
func Productos_delete(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])

	datos := modelos.Producto{}
	if err := database.Database.First(&datos, id); err.Error != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Recurso no disponible",
		}
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(respuesta)
		return
	} else {
		database.Database.Delete(&datos)
		respuesta := map[string]string{
			"estado":  "ok",
			"mensaje": "Se eliminó el registro exitosamente",
		}
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(respuesta)
	}
}
