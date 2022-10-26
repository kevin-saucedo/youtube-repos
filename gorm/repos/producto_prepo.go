package repos

import (
	"errors"

	"github.com/user0608/ytgorm/models"
	"gorm.io/gorm"
)

type ProductoRepository struct {
	conn *gorm.DB
}

func NewProductoRepository(c *gorm.DB) *ProductoRepository {
	return &ProductoRepository{conn: c}
}

func (r *ProductoRepository) Create(p *models.Producto) error {
	rs := r.conn.Create(p)
	if rs.Error != nil {
		return rs.Error
	}
	return nil
}

func (r *ProductoRepository) FindAll() ([]*models.Producto, error) {
	var productos []*models.Producto
	if err := r.conn.Find(&productos).Error; err != nil {
		return nil, err
	}
	return productos, nil
}
func (r *ProductoRepository) FindByCode(codigo string) (*models.Producto, error) {
	var producto models.Producto
	rs := r.conn.Find(&producto, codigo)
	if rs.Error != nil {
		return nil, rs.Error
	}
	if rs.RowsAffected == 0 {
		return nil, errors.New("registro no encontrado")
	}
	return &producto, nil
}
func (r *ProductoRepository) FindByName(nombre string) (*models.Producto, error) {
	var producto models.Producto
	rs := r.conn.Where("nombre =  ?", nombre).Limit(1).Find(&producto)
	if rs.Error != nil {
		return nil, rs.Error
	}
	if rs.RowsAffected == 0 {
		return nil, errors.New("registro no encontrado")
	}
	return &producto, nil
}

func (r *ProductoRepository) Delete(codigo string) error {
	rs := r.conn.Delete(&models.Producto{}, "codigo = ?", codigo)
	if rs.Error != nil {
		return rs.Error
	}
	if rs.RowsAffected == 0 {
		return errors.New("registro no encontrado")
	}
	return nil
}
func (r *ProductoRepository) DeleteWithQuery(codigo string) error {
	const sql = "delete from producto where codigo  = ?"
	rs := r.conn.Exec(sql, codigo)
	if rs.Error != nil {
		return rs.Error
	}
	if rs.RowsAffected == 0 {
		return errors.New("registro no encontrado")
	}
	return nil
}

func (r *ProductoRepository) FindAllWithQuery() ([]*models.Producto, error) {
	const qry = "select nombre,precio from producto"
	var productos []*models.Producto
	if err := r.conn.Raw(qry).Scan(&productos).Error; err != nil {
		return nil, err
	}
	return productos, nil
}
