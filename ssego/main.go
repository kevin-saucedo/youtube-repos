package main

import (
	"log"
	"net/http"
	"ssego/handlers"
)

func main() {
	r := http.NewServeMux()
	handlers.InitRoutes(r)
	log.Println("localhost:8080")
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Fatalln(err)
	}
}
