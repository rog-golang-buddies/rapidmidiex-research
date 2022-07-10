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
type clients map[uuid.UUID]*conn

func (cls clients) add(c *conn) {
	cls[c.id] = c
}

func (cls clients) broadcast(req websocket.Message) {
	for _, cli := range cls {
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
	// s.mu.Lock()
	// defer s.mu.Unlock()

	c := &conn{
		id:  uuid.New(),
		rwc: rwc,
		wsr: wsutil.NewReader(rwc, ws.StateServerSide),
		wsw: wsutil.NewWriter(rwc, ws.StateServerSide, ws.OpText),
		cls: s.r,
	}

	s.r[c.id] = c // connection has been made - tell everyone else in the pool
	return c
}

func (c *conn) serve() {
	defer c.rwc.Close()
	c.write(fmt.Sprintf("connection established: %s", c.id))

	for c.next() {
		var msg websocket.Message
		if c.err = c.read(&msg); c.err != nil {
			break
		}
		// send to clients
		go c.cls.broadcast(msg)
	}

	delete(c.cls, c.id)
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
