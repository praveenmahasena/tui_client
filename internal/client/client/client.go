package client

import (
	"bufio"
	"encoding/json"
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
	return &c, nil
}

func (c *Client) Start() error {

	con, conErr := net.Dial("tcp", ":42069")

	if conErr != nil {
		return conErr
	}

	go c.Write(con)
	//	go c.Read(con)

	cancelCh := make(chan os.Signal, 1)

	signal.Notify(cancelCh, syscall.SIGTERM, syscall.SIGINT)

	<-cancelCh

	return con.Close()
}

func (c Client) Read(con net.Conn) {

	r := bufio.NewReader(con)

	for {
		go func() {
			data, err := r.ReadBytes('\n')
			fmt.Println(string(data), err)
		}()
	}

}

func (c Client) Write(con net.Conn) {
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		c.Message = s.Text()
		json.NewEncoder(con).Encode(c)
	}
}
