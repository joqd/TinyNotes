package startup

import (
	"github.com/joqd/TinyNotes/internal/adapter/config"
	"github.com/joqd/TinyNotes/internal/adapter/server"
	"github.com/joqd/TinyNotes/internal/adapter/server/handler"
)

func LetsFuckingGo() {
	conf := config.Load()
	srv := server.New(conf)

	handlers := handler.Handlers{
		HealthHandler: handler.NewPingHandler(),
	}

	srv.SetupRouter(handlers)
	srv.Start()
}
