package routers

import (
	"encoding/json"
	"net/http"

	"github.com/GustGar16/twitter/bd"
	"github.com/GustGar16/twitter/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {
	var user models.Usuario
	//Se hace un decoder del body que viene en el request y se liga a la estructura User
	error := json.NewDecoder(r.Body).Decode(&user)
	if error != nil {
		http.Error(w, "Error en los parametros recibidos", 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "El email es requerido", 400)
		return
	}
	if len(user.Password) < 6 {
		http.Error(w, "El password debe ser superior a 6 caracteres", 400)
		return
	}

	//Verificar que no exista duplicidad de registros mediante email
	_, existe, _ := bd.UserVerify(user.Email)
	if existe == true {
		http.Error(w, "El correo ingresado ya existe", 400)
		return
	}

	_, status, err := bd.SuccessUserInsert(user)
	if err != nil {
		http.Error(w, "Error al intentar registrar el usuario", 400)
		return
	}
	if status == false {
		http.Error(w, "Error al almacenar usuario en base de datos", 400)
		return
	}
}
