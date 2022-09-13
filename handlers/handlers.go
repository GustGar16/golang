package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/GustGar16/twitter/middleware"
	"github.com/GustGar16/twitter/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {
	router := mux.NewRouter()

	//Declaracion de las rutas http
	//Primero se entrara al middleware DBVerifyConnection y una vez aprobado se redireccionará a la ruta indicada dentro de los parentesis
	router.HandleFunc("/registro", middleware.DBVerifyConnection(routers.CreateUser)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8088"
	}
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
