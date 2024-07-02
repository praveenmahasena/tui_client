package internal

import (
	"github.com/praveenmahasena/tui_chat_client/internal/publisher"
)

func Start() error {
	p, err := publisher.New()
	if err != nil {
		return err
	}
	return p.Run()
}
