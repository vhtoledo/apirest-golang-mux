package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)
type ResponseGnerico struct{
	Estado string
	Mensaje string
}
func Ejemplo_get(response http.ResponseWriter, request *http.Request) {
	
	response.Header().Set("COntent-Type", "application/json")
	response.Header().Add("victor", "www.victortoledodev.com.ar")
	output, _:= json.Marshal(ResponseGnerico{"ok", "Metodo GET"})
	fmt.Fprintln(response, string(output))
}

func Ejemplo_get_querystring(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("COntent-Type", "application/json")
	response.Header().Add("victor", "www.victortoledodev.com.ar")
	output, _:= json.Marshal(ResponseGnerico{"ok", "query string | id="+request.URL.Query().Get("id")})
	fmt.Fprintln(response, string(output))
}

func Ejemplo_get_con_parametros(response http.ResponseWriter, request *http.Request) {
	vars:=mux.Vars(request)
	response.Header().Set("COntent-Type", "application/json")
	response.Header().Add("victor", "www.victortoledodev.com.ar")
	output, _:= json.Marshal(ResponseGnerico{"ok", "Metodo GET | id =" + vars["id"]})
	fmt.Fprintln(response, string(output))
	
}
func Ejemplo_post(response http.ResponseWriter, request *http.Request){
	response.Header().Set("COntent-Type", "application/json")
	response.Header().Add("victor", "www.victortoledodev.com.ar")
	respuesta := map[string]string{
		"estado": "ok",
		"mensaje": "Metodo POST 2",
	}
	//response.WriteHeader(201)
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(respuesta)
}
/*
func Ejemplo_post(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("COntent-Type", "application/json")
	response.Header().Add("victor", "www.victortoledodev.com.ar")
	output, _:= json.Marshal(ResponseGnerico{"ok", "Metodo POST"})
	fmt.Fprintln(response, string(output))
}
*/
func Ejemplo_put(response http.ResponseWriter, request *http.Request) {
	vars:=mux.Vars(request)
	response.Header().Set("COntent-Type", "application/json")
	response.Header().Add("victor", "www.victortoledodev.com.ar")
	output, _:= json.Marshal(ResponseGnerico{"ok", "Metodo PUT| id =" + vars["id"]})
	fmt.Fprintln(response, string(output))
}

func Ejemplo_delete(response http.ResponseWriter, request *http.Request) {
	vars:=mux.Vars(request)
	response.Header().Set("COntent-Type", "application/json")
	response.Header().Add("victor", "www.victortoledodev.com.ar")
	output, _:= json.Marshal(ResponseGnerico{"ok", "Metodo DELETE | id =" + vars["id"]})
	fmt.Fprintln(response, string(output))
}