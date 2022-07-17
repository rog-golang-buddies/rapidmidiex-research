# Audio to MIDI using Audio API and Web Sockets

this is an example of note detection for instruments without MIDI support. it detects the notes you play and converts them to MIDI number of the note.

## How to run?

run Server:

- first install the API dependencies using `go mod tidy` command.
- start the server using `go run cmd/main.go` command.

run Client:

- move to client directory: `cd client`
- install node modules: `yarn` or `npm install`
- start client: `yarn run dev`
- open _http://localhost:3000_ in your browser

## What is an MIDI number?

the MIDI number of a note corresponds to the identifier number of the note in MIDI table.
for example the number 60 represents the note C4 which is the middle C on a piano.

more info here: [MIDI table](https://www.inspiredacoustics.com/en/MIDI_note_numbers_and_center_frequencies)

## What does the state field in the output mean?

state field shows if the note is playing or not (think NOTE_ON/NOTE_OFF in MIDI devices).
sending off state is not implemented yet.

### TODO

- [ ] detect multiple notes playing at the same time
- [ ] detect the velocity of playing note
- [ ] add cli support
- [ ] add midi to audio support
- [ ] add mult-client connection support
