package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthHandler interface {
	Ping(c echo.Context) error
}

type healthHandler struct{}

func NewPingHandler() HealthHandler {
	return &healthHandler{}
}

func (hh *healthHandler) Ping(c echo.Context) error {
	return c.Render(http.StatusOK, "ping", "Pong")
}
