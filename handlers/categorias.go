package handlers

import (
	"golang-mux-apirest/database"
	"golang-mux-apirest/dto"
	"golang-mux-apirest/modelos"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gosimple/slug"
)

func Categoria_get(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	datos := modelos.Categorias{}
	database.Database.Order("id desc").Find(&datos)
	//database.Database.Find(&datos)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(datos)
}
func Categoria_get_con_parametro(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	//id, _ = strconv.Atoi(vars["id"])
	datos := modelos.Categoria{}
	if err := database.Database.First(&datos, vars["id"]); err.Error != nil {
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
func Categoria_post(response http.ResponseWriter, request *http.Request) {
	var categoria dto.CategoriaDto
	if err := json.NewDecoder(request.Body).Decode(&categoria); err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(respuesta)
		return
	}
	datos := modelos.Categoria{Nombre: categoria.Nombre, Slug: slug.Make(categoria.Nombre)}
	database.Database.Save(&datos)
	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Se creó el registro exitosamente",
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(respuesta)
}
func Categoria_put(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])
	var categoria dto.CategoriaDto

	if err := json.NewDecoder(request.Body).Decode(&categoria); err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(respuesta)
		return
	}
	datos := modelos.Categoria{}
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
		datos.Nombre = categoria.Nombre
		datos.Slug = slug.Make(categoria.Nombre)
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
func Categoria_delete(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])

	datos := modelos.Categoria{}
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
