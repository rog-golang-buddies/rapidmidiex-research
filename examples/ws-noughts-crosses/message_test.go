package websocket

import (
	"encoding/json"
	"testing"
)

// test encoding for messageType
func TestMessageTypeMarshalJSON(t *testing.T) {
	var m MessageType
	m = Connected
	b, err := json.Marshal(m)
	if err != nil {
		t.Error(err)
	}
	if string(b) != "\"connected\"" {
		t.Errorf("expected \"connected\", got %s", string(b))
	}
}

// test decoding for messageType
func TestMessageTypeUnmarshalJSON(t *testing.T) {
	var m MessageType
	b := []byte("\"disconnected\"")
	err := json.Unmarshal(b, &m)
	if err != nil {
		t.Error(err)
	}
	if m != Disconnected {
		t.Errorf("expected disconnected, got %s", m)
	}
}

// test encoding for message
func TestMessageMarshalJSON(t *testing.T) {
	var m Message
	m.Type = Connected
	m.Data = struct {
		Id string `json:"id"`
	}{Id: "123"}
	b, err := json.Marshal(m)
	if err != nil {
		t.Error(err)
	}
	if string(b) != "{\"type\":\"connected\",\"data\":{\"id\":\"123\"}}" {
		t.Errorf("expected {\"type\":\"connected\",\"data\":{\"id\":\"123\"}}, got %s", string(b))
	}
}
