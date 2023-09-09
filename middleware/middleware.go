package middleware

import (
	"golang-mux-apirest/database"
	"golang-mux-apirest/modelos"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
	//jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func ValidarJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		//se valida existencia .env y variables de entorno
		errorVariables := godotenv.Load()
		if errorVariables != nil {
			respuesta := map[string]string{
				"estado":  "error",
				"mensaje": "Ocurrió un error inesperado",
			}
			response.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(response).Encode(respuesta)
			return
		}
		//empezamos a trabajar
		miClave := []byte(os.Getenv("SECRET_JWT"))
		response.Header().Add("content-type", "application/json")
		header := request.Header.Get("Authorization")
		//fmt.Println(len(token))
		if len(header) == 0 {
			respuesta := map[string]string{
				"estado":  "error",
				"mensaje": "No autorizado",
			}
			response.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(response).Encode(respuesta)
			return
		}
		splitBearer := strings.Split(header, " ")
		if len(splitBearer) != 2 {
			respuesta := map[string]string{
				"estado":  "error",
				"mensaje": "No autorizado",
			}
			response.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(response).Encode(respuesta)
			return
		}
		splitToken := strings.Split(splitBearer[1], ".")
		if len(splitToken) != 3 {
			respuesta := map[string]string{
				"estado":  "error",
				"mensaje": "No autorizado",
			}
			response.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(response).Encode(respuesta)
			return
		}
		tk := strings.TrimSpace(splitBearer[1])
		token, err := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: ")

			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return miClave, nil
		})
		//fmt.Println(splitBearer[1])
		if err != nil {
			respuesta := map[string]string{
				"estado":  "error",
				"mensaje": "No autorizado",
			}
			response.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(response).Encode(respuesta)
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			//fmt.Println(claims["id"])
			usuario := modelos.Usuario{}
			if err := database.Database.Where("correo = ?", claims["correo"]).First(&usuario); err.Error != nil {
				respuesta := map[string]string{
					"estado":  "error",
					"mensaje": "No autorizado",
				}
				response.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(response).Encode(respuesta)
				return
			} else {
				//fmt.Println(usuario.Nombre)
				next.ServeHTTP(response, request)
			}

		} else {
			respuesta := map[string]string{
				"estado":  "error",
				"mensaje": "No autorizado",
			}
			response.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(response).Encode(respuesta)
			return
		}

	}
}
