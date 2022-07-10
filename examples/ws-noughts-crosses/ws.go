package websocket

type contextKey struct {
	name string
}

func (k *contextKey) String() string { return "net/http context value " + k.name }

type MessageType int

const (
	Connected MessageType = iota
	Disconnected
	Join
	Leave
	Play
	// Unknown
)

func (m MessageType) String() string {
	switch m {
	case Connected:
		return "connected"
	case Disconnected:
		return "disconnected"
	case Join:
		return "join"
	case Leave:
		return "leave"
	case Play:
		return "play"
	default:
		return "unknown"
	}
}

type Message struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}
