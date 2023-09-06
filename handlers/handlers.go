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
	
	output, _:= json.Marshal(ResponseGnerico{"ok", "Metodo GET"})
	fmt.Fprintln(response, string(output))
}

func Ejemplo_get_querystring(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "query string | id="+request.URL.Query().Get("id"))
}

func Ejemplo_get_con_parametros(response http.ResponseWriter, request *http.Request) {
	vars:=mux.Vars(request)
	fmt.Fprintln(response, "hola ejemplo con parametros | id ="+vars["id"])
}

func Ejemplo_post(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Metodo post")
}

func Ejemplo_put(response http.ResponseWriter, request *http.Request) {
	vars:=mux.Vars(request)
	fmt.Fprintln(response, "Metodo put | id ="+vars["id"])
}

func Ejemplo_delete(response http.ResponseWriter, request *http.Request) {
	vars:=mux.Vars(request)
	fmt.Fprintln(response, "Metodo delete | id ="+vars["id"])
}