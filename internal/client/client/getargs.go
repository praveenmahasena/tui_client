package client

import (
	"github.com/charmbracelet/huh"
)

func (c *Client) GetArgs() error {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Options(huh.NewOptions("Public Chat", "Private Chat", "Group Chat")...).
				Title("Chat Type").
				Value(&c.ChatType),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("User Name").
				Value(&c.UserName),
		),
	)
	return form.Run()
}
