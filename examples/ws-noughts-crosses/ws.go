package websocket

type contextKey struct {
	name string
}

func (k *contextKey) String() string { return "net/http context value " + k.name }

type Message struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

type Play struct {
	Turn  int
	Pos   int
	Moves int
}
