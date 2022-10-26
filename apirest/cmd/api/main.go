package main

import (
	"apirest/cmd/api/routes"
	"apirest/connection"
	"log"
	"net/http"
)

const (
	host     = "127.0.0.1"
	port     = "5432"
	user     = "postgres"
	password = "maira002"
	dbname   = "prueba_db"
)

func main() {
	if err := connection.InitConn(host, port, user, password, dbname); err != nil {
		log.Fatalln("DB Err:", err)
	}
	log.Println("DB Status.....OK")
	mux := http.NewServeMux()
	routes.Routes(mux)
	log.Println(":8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalln(err)
	}
}
