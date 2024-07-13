package internal

import "github.com/praveenmahasena/tui_chat_client/internal/client/client"

func Start() error {
	c, err := client.New()
	if err != nil {
		return err
	}

	return c.Start()
}
