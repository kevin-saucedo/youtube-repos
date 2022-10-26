package services

import (
	"errors"

	"github.com/user0608/ytgorm/models"
	"github.com/user0608/ytgorm/repos"
)

type ProductoService struct {
	repo *repos.ProductoRepository
}

func NewProductoService(r *repos.ProductoRepository) *ProductoService {
	return &ProductoService{repo: r}
}

func (s *ProductoService) Registrar(p *models.Producto) error {
	//validar data
	return s.repo.Create(p)
}
func (s *ProductoService) GetAll() ([]*models.Producto, error) {
	return s.repo.FindAll()
}
func (s *ProductoService) GetByCodigo(codigo string) (*models.Producto, error) {
	if codigo == "" {
		return nil, errors.New("codigo invalido")
	}
	return s.repo.FindByCode(codigo)
}
func (s *ProductoService) Delete(codigo string) error {
	if codigo == "" {
		return errors.New("codigo invalido")
	}
	return s.repo.DeleteWithQuery(codigo)
}
