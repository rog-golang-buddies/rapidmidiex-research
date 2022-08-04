package ws

import (
	"fmt"
	"sync"

	chat "ws.rog.noughtscrosses"
)

type Pool struct {
	mu sync.RWMutex

	r   chan *Client
	unr chan *Client
	bc  chan chat.Message

	clis map[*Client]bool
}

func NewPool() *Pool {
	p := &Pool{
		r:    make(chan *Client),
		unr:  make(chan *Client),
		bc:   make(chan chat.Message),
		clis: make(map[*Client]bool),
	}
	go p.run()
	return p
}

func (p *Pool) broadcast(v any) {
	for cli := range p.clis {
		cli.write(v)
	}
}

func (p *Pool) add(cli *Client) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.clis[cli] = true
	// This will give an error in the client as it does not conform to my noughtscrosses.Message type
	p.broadcast(fmt.Sprintf("New User Joined... Size of Connection Pool: %d", len(p.clis)))

	go cli.serve()
}

func (p *Pool) remove(cli *Client) {
	p.mu.Lock()
	defer func() {
		p.mu.Unlock()
		cli.close()
	}()

	delete(p.clis, cli)
	// This will give an error in the client as it does not conform to my noughtscrosses.Message type
	p.broadcast(fmt.Sprintf("User Disconnected... Size of Connection Pool: %d", len(p.clis)))
}

func (p *Pool) run() {
	for {
		select {
		case cli := <-p.r:
			p.add(cli)
		case cli := <-p.unr:
			p.remove(cli)
		case msg := <-p.bc:
			p.broadcast(msg)
		}
	}
}
