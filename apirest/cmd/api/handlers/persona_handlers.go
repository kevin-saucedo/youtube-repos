package handlers

import (
	"apirest/models"
	"apirest/services"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func FindAllPersons(service services.PersonaService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		personas, err := service.GetAll()
		if err != nil {
			log.Println("handlers:FindAllPersons:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(personas); err != nil {
			log.Println("handlers:FindAllPersons:json_encoding:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func CreatePerson(service services.PersonaService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var persona models.Persona
		if err := json.NewDecoder(r.Body).Decode(&persona); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := service.Registrar(persona); err != nil {
			if errors.Is(err, services.ErrEmptyNombre) || errors.Is(err, services.ErrEmptyApellido) {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			log.Println("handlers:CreatePerson:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
