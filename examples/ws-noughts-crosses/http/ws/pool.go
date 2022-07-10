package ws

import (
	"net"
	"sync"

	websocket "github.com/rog-golang-buddies/realtime-midi/examples/ws-noughts-crosses"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/google/uuid"
)

type Pool struct {
	// Read/Write mutex so that we can safely read/write from/to the map
	mu sync.RWMutex
	m  map[uuid.UUID]*Conn
}

func NewPool() *Pool {
	return &Pool{m: make(map[uuid.UUID]*Conn)}
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
	p.mu.Lock()
	defer p.mu.Unlock()

	p.m[c.id] = c
}

func (p *Pool) delete(c *Conn) {
	p.mu.Lock()
	defer p.mu.Unlock()

	delete(p.m, c.id)
}

// Send message to all connections in pool
func (p *Pool) Broadcast(req websocket.Message) {
	for _, cli := range p.m {
		cli.Write(req)
	}
}
