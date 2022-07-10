package websocket

import "encoding/json"

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
	Unknown
)

// custom marshaler for MessageType
func (m MessageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}

// custom unmarshaler for MessageType
func (m *MessageType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch s {
	case "connected":
		*m = Connected
	case "disconnected":
		*m = Disconnected
	case "join":
		*m = Join
	case "leave":
		*m = Leave
	case "play":
		*m = Play
	default:
		*m = Unknown
	}

	return nil
}

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
	Type MessageType `json:"type"`
	Data any         `json:"data"`
}
