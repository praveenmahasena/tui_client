package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
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

	go c.write(con)
	go c.read(con)

	cancelCh := make(chan os.Signal, 1)

	signal.Notify(cancelCh, syscall.SIGTERM, syscall.SIGINT)

	<-cancelCh

	return con.Close()
}

func (c *Client) read(con net.Conn) {
	prefix := c.makeHeaderMsg()

	scanner := bufio.NewScanner(con)

	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, prefix) {
			continue
		}
		fmt.Println(text)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner has an error: %s\n", err)
	}
}

func (c *Client) write(con net.Conn) {
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		_, err := con.Write([]byte(fmt.Sprintf("%s: %s\n", c.makeHeaderMsg(), s.Text())))
		if err != nil {
			fmt.Printf("unable to write: %s\n", err)
		}
	}
}

func (c *Client) makeHeaderMsg() string {
	return fmt.Sprintf("Message From %s", c.UserName)
}
