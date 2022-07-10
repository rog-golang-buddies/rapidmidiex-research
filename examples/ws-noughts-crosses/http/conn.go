package http

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"

	"ws.rog.numberguesser"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/google/uuid"
)

// basically a room
type clients struct {
	m map[uuid.UUID]*conn
}

var defaultClients = &clients{m: make(map[uuid.UUID]*conn)}

// func (cls clients) add(c *conn) {
// 	cls[c.id] = c
// }

func (cls clients) broadcast(req websocket.Message) {
	for i, cli := range cls.m {
		log.Println(">", i)
		cli.write(req)
	}
}

type conn struct {
	id  uuid.UUID
	rwc net.Conn
	wsr *wsutil.Reader
	wsw *wsutil.Writer
	err error
	cls clients
}

func (s *Service) newConn(rwc net.Conn) *conn {
	c := &conn{
		id:  uuid.New(),
		rwc: rwc,
		wsr: wsutil.NewReader(rwc, ws.StateServerSide),
		wsw: wsutil.NewWriter(rwc, ws.StateServerSide, ws.OpText),
		cls: s.cls,
	}

	s.cls.m[c.id] = c // connection has been made - tell everyone else in the pool
	return c
}

func (c *conn) serve() {
	defer c.rwc.Close()

	message := struct {
		Type string `json:"type"`
		Data struct {
			Id string `json:"id"`
		} `json:"data"`
	}{Type: "join", Data: struct {
		Id string "json:\"id\""
	}{Id: c.id.String()}}

	c.write(message)

	log.Println(c.cls.m)

	for c.next() {
		var msg websocket.Message
		if c.err = c.read(&msg); c.err != nil {
			break
		}
		fmt.Println(len(c.cls.m))
		// send to clients
		go c.cls.broadcast(msg)
	}

	delete(c.cls.m, c.id)
	log.Printf("connection closed: %s", c.err)
}

func (c *conn) next() bool {
	h, err := c.wsr.NextFrame()
	if err != nil {
		c.err = err
		return false
	}

	if h.OpCode == ws.OpClose {
		c.err = io.EOF
		return false
	}
	return true //continue as no error
}

func (c *conn) read(v any) error { return json.NewDecoder(c.wsr).Decode(v) }

func (c *conn) write(v any) error {
	if err := json.NewEncoder(c.wsw).Encode(v); err != nil {
		return err
	}
	return c.wsw.Flush()
}
