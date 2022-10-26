package models

type Producto struct {
	Codigo string  `json:"codigo" gorm:"primaryKey"`
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
}
