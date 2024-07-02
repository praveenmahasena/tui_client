package publisher

import (
	"os"

	"github.com/charmbracelet/log"
)

func handleErr(err error) {
	l := log.New(os.Stdout)
	l.Warn(err)
}
