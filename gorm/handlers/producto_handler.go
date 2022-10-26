package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/user0608/ytgorm/models"
	"github.com/user0608/ytgorm/services"
)

type ProductoHandler struct {
	binder  echo.DefaultBinder
	service *services.ProductoService
}

func NewProductoHandler(s *services.ProductoService) *ProductoHandler {
	return &ProductoHandler{
		service: s,
		binder:  echo.DefaultBinder{},
	}
}

func (h *ProductoHandler) CreateProducto(c echo.Context) error {
	var producto models.Producto
	if err := h.binder.BindBody(c, &producto); err != nil {
		return c.JSON(http.StatusBadRequest, Response{Message: "json invalido"})
	}
	if err := h.service.Registrar(&producto); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, Response{Message: "algo paso"})
	}
	return c.JSON(http.StatusOK, Response{Message: "OK,Registro Guardado!"})
}

func (h *ProductoHandler) FindAll(c echo.Context) error {
	productos, err := h.service.GetAll()
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, Response{Message: "algo paso"})
	}
	return c.JSON(http.StatusOK, Response{Data: productos})
}
func (h *ProductoHandler) FindByCode(c echo.Context) error {
	producto, err := h.service.GetByCodigo(c.Param("producto_codigo"))
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, Response{Message: "algo paso"})
	}
	return c.JSON(http.StatusOK, Response{Data: producto})
}
func (h *ProductoHandler) DeleteProducto(c echo.Context) error {
	if err := h.service.Delete(c.Param("producto_codigo")); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, Response{Message: "algo paso"})
	}
	return c.JSON(http.StatusOK, Response{Message: "Ok!!!!!!!"})
}
