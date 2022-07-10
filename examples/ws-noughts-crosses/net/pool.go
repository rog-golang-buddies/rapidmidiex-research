package net

import (
	"log"
	"net"

	"ws.rog.numberguesser"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/google/uuid"
)

type Pool struct {
	m map[uuid.UUID]*Conn // in-order to overcome race condition, this needs to be a
}

func NewPool() *Pool {
	return &Pool{make(map[uuid.UUID]*Conn)}
}

// Create a new connection and keep on an infinite loop
func (p *Pool) NewConn(rwc net.Conn) *Conn {
	c := &Conn{
		id:   uuid.New(),
		rwc:  rwc,
		wsr:  wsutil.NewReader(rwc, ws.StateServerSide),
		wsw:  wsutil.NewWriter(rwc, ws.StateServerSide, ws.OpText),
		pool: p,
	}
	p.add(c)

	go c.serve()
	return c
}

func (p *Pool) add(c *Conn) {
	p.m[c.id] = c
}

func (p *Pool) delete(c *Conn) {
	delete(p.m, c.id)
}

// Send message to all connections in pool
func (p Pool) Broadcast(req websocket.Message) {
	for i, cli := range p.m {
		log.Println(">", i)
		cli.Write(req)
	}
}
