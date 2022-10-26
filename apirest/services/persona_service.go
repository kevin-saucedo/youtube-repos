package services

import (
	"apirest/connection"
	"apirest/models"
	"apirest/repositories"
	"errors"
)

var (
	ErrEmptyNombre   = errors.New("error nombre persona vacio")
	ErrEmptyApellido = errors.New("error apellido persona vacio")
)

type PersonaService interface {
	GetAll() ([]models.Persona, error)
	Registrar(p models.Persona) error
}
type persona struct {
	r repositories.PersonaRepository
}

func NewPersonaService(r repositories.PersonaRepository) PersonaService {
	return &persona{r}
}

func (s *persona) GetAll() ([]models.Persona, error) {
	return s.r.FindAll(connection.Conn)
}
func (s *persona) Registrar(p models.Persona) error {
	if p.Nombre == "" {
		return ErrEmptyNombre
	}
	if p.Apellido == "" {
		return ErrEmptyApellido
	}
	return s.r.Create(connection.Conn, p)
}
