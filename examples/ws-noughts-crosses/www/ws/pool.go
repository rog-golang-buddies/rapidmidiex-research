package ws

import (
	"fmt"
	"sync"

	chat "ws.rog.noughtscrosses"
)

type Pool struct {
	mu sync.RWMutex

	register   chan *Client
	unregister chan *Client
	messages   chan chat.Message

	clients map[*Client]bool
}

func NewPool() *Pool {
	p := &Pool{
		register:   make(chan *Client),
		unregister: make(chan *Client),
		messages:   make(chan chat.Message),
		clients:    make(map[*Client]bool),
	}
	go p.run()
	return p
}

func (p *Pool) broadcast(v any) {
	for cli := range p.clients {
		cli.write(v)
	}
}

func (p *Pool) add(cli *Client) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.clients[cli] = true

	p.broadcast(fmt.Sprintf("New User Joined... Size of Connection Pool: %d", len(p.clients)))

	go cli.serve()
}

func (p *Pool) remove(cli *Client) {
	p.mu.Lock()
	defer func() {
		p.mu.Unlock()
		cli.close()
	}()

	delete(p.clients, cli)
	// This will give an error in the client as it does not conform to my noughtscrosses.Message type
	p.broadcast(fmt.Sprintf("User Disconnected... Size of Connection Pool: %d", len(p.clients)))
}

func (p *Pool) run() {
	for {
		select {
		case cli := <-p.register:
			p.add(cli)
		case cli := <-p.unregister:
			p.remove(cli)
		case msg := <-p.messages:
			p.broadcast(msg)
		}
	}
}
