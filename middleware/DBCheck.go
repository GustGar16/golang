package middleware

import (
	"net/http"

	"github.com/GustGar16/twitter/bd"
)

/*
Hace la verificacion de la conexion a la base datos
Para un middleware se debe recibir y retornar una fucion de tipo http
*/
func DBVerifyConnection(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Si no existe conexion se detiene todo y se arroja un error 500
		if bd.CheckBD() == 0 {
			http.Error(w, "Conexion perdida en la base de datos (mid [DBVerifyConnection])", 500)
			return
		}
		//Si existe conexion a la BD se da paso y se retorna el http original para su uso correcto
		next.ServeHTTP(w, r)
	}
}
