# Audio to MIDI using Audio API and Web Sockets

this is an example of note detection for instruments without MIDI support. it detects the notes you play and converts them to MIDI number of the note.

## How to run?

run Server:

- first install the API dependencies using _go mod tidy_ command.
- start the server using _go run cmd/main.go_ command.

run Client:

- move to client directory: _cd client_
- install node modules: _yarn_ or _npm install_
- start client: _yarn run dev_
- open _http://localhost:3000_ in your browser

## What is an MIDI number?

the MIDI number of a note corresponds to the identifier number of the note in MIDI table.
for example the number 60 represents the note C4 which is the middle C on a piano.

more info here: [MIDI table](https://www.inspiredacoustics.com/en/MIDI_note_numbers_and_center_frequencies)

## What does the state field in the output mean?

state field shows if the note is playing or not (think NOTE_ON/NOTE_OFF in MIDI devices).
sending off state is not implemented yet.

### TODO

- detect multiple notes playing at the same time
- detect the velocity of playing note
- add cli support
- add midi to audio support
- add mult-client connection support
