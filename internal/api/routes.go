package api

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func (s *Server) RegisterRoutes(e *echo.Echo) {
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/:code", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Ooi")
	})

	// v1 := e.Group("/api/v1")
	// v1.POST()
	// v1.GET()

	e.GET("/health", s.handleHealth)
}
