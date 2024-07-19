package dialer

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/praveenmahasena/tui_chat_client/internal/user"
)

type Dialer struct {
	Port    string
	Network string
	Ctx     context.Context
}

func New(p, n string) *Dialer {
	return &Dialer{
		Port:    p,
		Network: n,
	}
}

func (d *Dialer) Dial(u *user.User) error {
	con, conErr := net.Dial(d.Network, d.Port)

	if conErr != nil {
		return conErr
	}

	go shutdown(con)
	handle(u, con)
	return nil
}

func handle(u *user.User, con net.Conn) {

	go read(con)
	write(con, u)

}

func read(con net.Conn) {

	s := bufio.NewScanner(con)

	for s.Scan() {

		if s.Err() != nil {
			log.Println(s.Err())
			continue
		}

		fmt.Println(s.Text())

	}

}

func write(con net.Conn, u *user.User) {

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		if _, err := con.Write([]byte("from" + u.Name + s.Text() + "\n")); err != nil {
			log.Println(err)
		}
	}

}

func shutdown(c net.Conn) {
	cancelCh := make(chan os.Signal, 1)

	signal.Notify(cancelCh, os.Interrupt, syscall.SIGTERM)

	<-cancelCh

	c.Write([]byte("closing"))
	c.Close()

}
