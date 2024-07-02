package publisher

import "github.com/charmbracelet/huh"

func GetArgs(pub *Pub) (*Pub, error) {

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				CharLimit(15).
				Title("Chatter Name").
				Value(&pub.User),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Options(
					huh.NewOptions("Group Chat", "Private Chat", "Public Chat")...,
				).
				Title("Chat Type You Wish to Join").
				Value(&pub.ChatType),
		),
	)

	if err := form.Run(); err != nil {
		return nil, err
	}

	return pub, nil
}
