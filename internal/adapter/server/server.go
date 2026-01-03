package server

import (
	"fmt"

	"github.com/joqd/TinyNotes/internal/adapter/config"
	"github.com/labstack/echo/v4"
)

type Server struct {
	e    *echo.Echo
	conf *config.Config
}

func New(conf *config.Config) *Server {
	e := echo.New()
	e.HideBanner = true

	e.Debug = conf.App.Debug

	return &Server{e: e, conf: conf}
}

func (s *Server) Start() error {
	address := fmt.Sprintf(":%d", s.conf.App.Port)
	return s.e.Start(address)
}

func (s *Server) GetEcho() *echo.Echo {
	return s.e
}
