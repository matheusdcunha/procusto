package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
)

var version string = "1.0.0"

func (s *Server) handleHealth(c *echo.Context) error {

	type health struct {
		Status    string    `json:"status"`
		Version   string    `json:"version"`
		Timestamp time.Time `json:"timestamp"`
	}

	h := health{"available", version, time.Now()}

	return c.JSON(http.StatusOK, h)
}
