package routes

import (
	"apirest/cmd/api/handlers"
	"apirest/repositories"
	"apirest/services"
	"net/http"
)

func Routes(mux *http.ServeMux) {
	repo := repositories.NewPersonaRepository()
	service := services.NewPersonaService(repo)
	mux.HandleFunc("/get", handlers.FindAllPersons(service))
	mux.HandleFunc("/create", handlers.CreatePerson(service))
}
