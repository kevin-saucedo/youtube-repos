package repositories

import (
	"apirest/models"
	"database/sql"
)

type PersonaRepository interface {
	FindAll(tx *sql.DB) ([]models.Persona, error)
	Create(tx *sql.DB, p models.Persona) error
}
type persona struct{}

func NewPersonaRepository() PersonaRepository {
	return &persona{}
}

func (r *persona) FindAll(tx *sql.DB) ([]models.Persona, error) {
	const qry = "select id,nombre,apellido,direccion,telefono from  persona"
	rows, err := tx.Query(qry)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var personas []models.Persona
	for rows.Next() {
		var p models.Persona
		err = rows.Scan(&p.ID, &p.Nombre, &p.Apellido, &p.Direccion, &p.Telefono)
		if err != nil {
			return nil, err
		}
		personas = append(personas, p)
	}
	return personas, nil
}
func (r *persona) Create(tx *sql.DB, p models.Persona) error {
	const qry = "insert into persona(nombre,apellido,direccion,telefono) values($1,$2,$3,$4)"
	_, err := tx.Exec(qry, p.Nombre, p.Apellido, p.Direccion, p.Telefono)
	if err != nil {
		return err
	}
	return nil
}
