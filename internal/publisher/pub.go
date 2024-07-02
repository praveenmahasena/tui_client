package publisher

import (
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Pub struct {
	User     string
	ChatType string
}

func New() (*Pub, error) {
	p := Pub{}

	return GetArgs(&p)
}

func (p *Pub) Run() error {
	con, conErr := net.Dial("tcp", ":42069")

	if conErr != nil {
		return conErr
	}

	//ctx, cancel := context.WithCancel(context.Background())

	go handle(con)
	// I could've used ctx here but meh

	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, syscall.SIGKILL, syscall.SIGINT)

	<-sigCh

	con.Write([]byte("Left"))
	return con.Close()
}
