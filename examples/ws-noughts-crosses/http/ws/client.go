package ws

import (
	"fmt"

	"github.com/gorilla/websocket"

	chat "ws.rog.noughtscrosses"
)

func NewClient(p *Pool, rwc *websocket.Conn) error {
	c := &Client{rwc, p}
	p.r <- c
	return c.serve()
}

type Client struct {
	rwc *websocket.Conn
	p   *Pool
}

func (c *Client) close() error {
	c.p.unr <- c
	return c.rwc.Close()
}

func (cli Client) read(v any) error { return cli.rwc.ReadJSON(v) }

func (cli Client) write(v any) error { return cli.rwc.WriteJSON(v) }

// TODO -- if error, close connection
func (cli *Client) serve() error {

	for {
		var msg chat.Message
		if err := cli.read(&msg); err != nil {
			cli.p.unr <- cli
			return err
		}

		// doSomething with message

		cli.p.bc <- msg
		fmt.Printf("Message Received: %+v\n", msg)
	}
}
