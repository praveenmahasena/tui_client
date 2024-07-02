package main

import (
	"log"

	"github.com/praveenmahasena/tui_chat_client/internal"
)

func main() {
	if err := internal.Start(); err != nil {
		log.Fatalln(err)
	}
}
