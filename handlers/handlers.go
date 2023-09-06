package handlers

import (
	"fmt"
	"net/http"
)

func Ejemplo_get(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "hola ejemplo")
}

func Ejemplo_get_con_parametros(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "hola ejemplo con parametros")
}

func Ejemplo_post(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Metodo post")
}

func Ejemplo_put(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Metodo put")
}

func Ejemplo_delete(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Metodo delete")
}