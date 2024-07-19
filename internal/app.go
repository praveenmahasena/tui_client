package internal

import (
	"github.com/praveenmahasena/tui_chat_client/internal/args"
	"github.com/praveenmahasena/tui_chat_client/internal/dialer"
	"github.com/praveenmahasena/tui_chat_client/internal/user"
)

func Start() error {
	u := user.New()

	if err := args.GetArgs(u); err != nil {
		return err
	}

	d := dialer.New(":42069", "tcp")

	return d.Dial(u)
}
