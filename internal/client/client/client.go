package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Client struct {
	UserName string
	ChatType string
	// Password string
	Message string
}

func New() (*Client, error) {
	c := Client{}
	if err := c.GetArgs(); err != nil {
		return nil, err
	}
	c.ChatType = "Public Chat"
	return &c, nil
}

func (c *Client) Start() error {

	con, conErr := net.Dial("tcp", ":42069")

	if conErr != nil {
		return conErr
	}

	go write(con)
	go read(con)

	cancelCh := make(chan os.Signal, 1)

	signal.Notify(cancelCh, syscall.SIGTERM, syscall.SIGINT)

	<-cancelCh

	return con.Close()
}

func read(con net.Conn) {

}

func write(con net.Conn) {
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		go func(c net.Conn) {
			n, err := c.Write(s.Bytes())
			fmt.Println(n, err)
		}(con)
	}
}
