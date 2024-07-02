package publisher

import (
	"bufio"
	"fmt"
	"net"
	"sync"

	"github.com/charmbracelet/huh"
)

func handle(con net.Conn) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	defer wg.Done()
	go read(con)
	go write(con)
	wg.Wait()
}

func read(c net.Conn) {
	scanner := bufio.NewScanner(c)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func write(c net.Conn) {

	for {
		var str string
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewText().CharLimit(100).Value(&str).Title("Write Your msg"),
			),
		)

		if err := form.Run(); err != nil {
			handleErr(err)
			continue
		}
		fmt.Fprintln(c, str)
	}

}
