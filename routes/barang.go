package routes

import (
	"backend/handlers"
	"backend/pkg/postgresql"
	"backend/repositories"

	"github.com/labstack/echo/v4"
)

func BarangRoutes(e *echo.Group) {
	barangRepository := repositories.RepositoryBarang(postgresql.DB)
	h := handlers.HandlerBarang(barangRepository)

	// e.POST("/barang", h.AddBarang)
	e.POST("/barang", h.AddBarang)
	e.GET("/barang/:id", h.GetBarang)
	e.GET("/barang", h.GetAllBarang)
	e.DELETE("/barang/:id", h.DeleteBarang)
	e.PUT("/barang/:id", h.UpdateBarang)
}
