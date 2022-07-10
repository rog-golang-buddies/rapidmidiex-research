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

type pool struct {
	m map[uuid.UUID]*conn // in-order to overcome race condition, this needs to be a
}

func (p *pool) add(c *conn) {
	p.m[c.id] = c
}

func (p *pool) delete(c *conn) {
	delete(p.m, c.id)
}

// Send message to all connections in pool
func (p pool) broadcast(req websocket.Message) {
	for i, cli := range p.m {
		log.Println(">", i)
		cli.write(req)
	}
}

type conn struct {
	id   uuid.UUID
	rwc  net.Conn
	wsr  *wsutil.Reader
	wsw  *wsutil.Writer
	err  error
	pool *pool
}

func (s *Service) newConn(rwc net.Conn) *conn {
	c := &conn{
		id:   uuid.New(),
		rwc:  rwc,
		wsr:  wsutil.NewReader(rwc, ws.StateServerSide),
		wsw:  wsutil.NewWriter(rwc, ws.StateServerSide, ws.OpText),
		pool: s.pool,
	}

	s.pool.add(c)
	return c
}

// Keep-alive connection using an inifinte for-loop.
// Once error, delete from pool & close connection
func (c *conn) serve() {
	defer c.rwc.Close()

	// anon type for "connected" message type, should move to root of package
	message := struct {
		Type string `json:"type"`
		Data struct {
			Id string `json:"id"`
		} `json:"data"`
	}{Type: "join", Data: struct {
		Id string "json:\"id\""
	}{Id: c.id.String()}}

	c.write(message)

	for c.next() {
		var msg websocket.Message
		if c.err = c.read(&msg); c.err != nil {
			break
		}
		fmt.Println(len(c.pool.m))
		// send to clients
		go c.pool.broadcast(msg)
	}

	c.pool.delete(c)
	log.Printf("connection closed: %s", c.err)
}

// Wrapper around the wsutil.Reader.NextFrame() method.
// Returns a bool type meaning can be used with an infinite for-loop
//
// If error preparing or EOF then returns false, breaking the for-loop
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

// Reads input from client
func (c *conn) read(v any) error { return json.NewDecoder(c.wsr).Decode(v) }

// Writes back to the client & flushes buffer onced it's finished
func (c *conn) write(v any) error {
	if err := json.NewEncoder(c.wsw).Encode(v); err != nil {
		return err
	}
	return c.wsw.Flush()
}
