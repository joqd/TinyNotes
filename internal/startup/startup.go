package startup

import (
	"log"

	"github.com/joqd/TinyNotes/internal/adapter/config"
)

func LetsFuckingGo() {
	conf := config.Load()
	log.Printf("app running on %d", conf.App.Port)
}
