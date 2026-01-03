package server

import "github.com/joqd/TinyNotes/internal/adapter/server/handler"

func (s *Server) SetupRouter(handlers handler.Handlers) {
	// Health Check
	healthGroup := s.e.Group("/health")
	healthGroup.GET("/ping", handlers.HealthHandler.Ping)
}
