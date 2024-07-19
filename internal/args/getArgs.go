package args

import (
	"github.com/praveenmahasena/tui_chat_client/internal/user"

	"github.com/charmbracelet/huh"
)

func GetArgs(u *user.User) error {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("user Name").
				Value(&u.Name),
		),
	)
	return form.Run()
}
