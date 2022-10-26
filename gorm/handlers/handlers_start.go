package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/user0608/ytgorm/repos"
	"github.com/user0608/ytgorm/services"
	"gorm.io/gorm"
)

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Start(e *echo.Echo, conn *gorm.DB) {
	productoHandler(e, conn)
}
func productoHandler(e *echo.Echo, conn *gorm.DB) {
	g := e.Group("/producto")
	h := NewProductoHandler(services.NewProductoService(repos.NewProductoRepository(conn)))
	g.POST("", h.CreateProducto)
	g.GET("", h.FindAll)
	g.GET("/:producto_codigo", h.FindByCode)
	g.DELETE("/:producto_codigo", h.DeleteProducto)
}
