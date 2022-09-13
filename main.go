package main

import (
	"log"

	"github.com/GustGar16/twitter/bd"
	"github.com/GustGar16/twitter/handlers"
)

func main() {
	if bd.CheckBD() == 0 {
		log.Fatal("Sin conexión a la BD (main)")
		return
	}
	handlers.Manejadores()
}
