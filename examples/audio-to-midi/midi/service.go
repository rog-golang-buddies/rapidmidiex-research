package midi

type midiSignal struct {
	MIDINum int  `json:"midi_num"` // number of the note in midi table (A4 is 69)
	State   bool `json:"state"`    // state of the note (on/off)
}

// func (s *midiSignal) getMIDINum() int {
// 	return s.MIDINum
// }

// func (s *midiSignal) getState() bool {
// 	return s.State
// }
