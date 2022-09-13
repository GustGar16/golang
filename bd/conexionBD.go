package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Connection = DBConnection()
var clientOptions = options.Client().ApplyURI("mongodb+srv://test:test@testtwitter.8hebryt.mongodb.net/?retryWrites=true&w=majority")

func DBConnection() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Error en la conexión a la BD")
		return client
	}
	//Hacemos un ping a la bd para comprobar la conexion a esta
	ping := client.Ping(context.TODO(), nil)
	if ping != nil {
		log.Printf(ping.Error())
		return client
	}
	log.Printf("Conexión exitosa a la bd")
	return client
}

func CheckBD() int {
	ping := Connection.Ping(context.TODO(), nil)
	if ping != nil {
		return 0
	}
	return 1
}
