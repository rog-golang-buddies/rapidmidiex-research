package websocket

import (
	"encoding/json"
	"testing"
)

// test encoding for messageType
func TestMessageTypeMarshalJSON(t *testing.T) {
	tests := []struct {
		m    MessageType
		want string
	}{
		{Connected, `"connected"`},
		{Disconnected, `"disconnected"`},
		{Join, `"join"`},
		{Leave, `"leave"`},
		{Play, `"play"`},
		{Reset, `"reset"`},
		{Unknown, `"unknown"`},
	}

	for _, test := range tests {
		b, err := json.Marshal(test.m)
		if err != nil {
			t.Error(err)
		}
		if string(b) != test.want {
			t.Errorf("\nexpected: [%s]\ngot:      [%s]", test.want, string(b))
		}
	}
}

// test decoding for messageType
func TestMessageTypeUnmarshalJSON(t *testing.T) {
	var m MessageType
	b := []byte(`"disconnected"`)
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

	expected := `{"type":"connected","data":{"id":"123"}}`

	b, err := json.Marshal(m)
	if err != nil {
		t.Error(err)
	}
	if string(b) != expected {
		t.Errorf("\nexpected: [%s]\ngot:      [%s]", expected, string(b))
	}
}
