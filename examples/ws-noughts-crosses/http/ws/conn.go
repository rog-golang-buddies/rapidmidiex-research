package ws

import (
	"encoding/json"
	"io"
	"log"
	"net"

	"ws.rog.noughtscrosses"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/google/uuid"
)

type Conn struct {
	id   uuid.UUID
	rwc  net.Conn
	wsr  *wsutil.Reader
	wsw  *wsutil.Writer
	err  error
	pool *Pool
}

// Keep-alive connection using an inifinte for-loop.
// Once error, delete from pool & close connection
func (c *Conn) serve() {
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

	c.Write(message)

	for c.next() {
		var msg websocket.Message
		if c.err = c.Read(&msg); c.err != nil {
			break
		}

		// doSomething with the msg

		// send to clients
		go c.pool.Broadcast(msg)
	}

	c.pool.delete(c)
	log.Printf("connection closed: %s", c.err)
}

// Wrapper around the wsutil.Reader.NextFrame() method.
// Returns a bool type meaning can be used with an infinite for-loop
//
// If error preparing or EOF then returns false, breaking the for-loop
func (c *Conn) next() bool {
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
func (c *Conn) Read(v any) error { return json.NewDecoder(c.wsr).Decode(v) }

// Writes back to the client & flushes buffer onced it's finished
func (c *Conn) Write(v any) error {
	if err := json.NewEncoder(c.wsw).Encode(v); err != nil {
		return err
	}
	return c.wsw.Flush()
}
