# Introduction

A study on how to let multiple musicians jam together (musically) over the Internet.

## Structure of this repo

Our own created diagrams go in [Diagrams](docs/diagrams/README.md).

Research and links to potentially useful libraries or examples go in [Research](docs/research/README.md).

## What does this project solve?

One of the most important elements of music is Time. For musicians to play together, they need to sync with each other. Even a slight delay between the sound you hear and the voice of the instrument significantly impacts the music's integrity and harmony. 
We can't always play together in person. Thereby we need to use online music jam services. The problem with these services is that they use the audio captured from your device and send it to other users who are listening to the music. Uploading audio data is so expensive that any minor instability in your internet connection has an enormous impact on the collaborative experience.

So we decided to look at this problem with a different POV. Why not send MIDI notes instead of audio? If you know some music theory then you know that the note A4 ("La" of the 4th octave in spanish) sounds like A4 everywhere. This way we can play notes in MIDI format as audio on other users' devices.

Here is an example:

1. User A plays note C on their MIDI device (a MIDI device can be a piano, a guitar or any instrument which has a MIDI port)<br/>
  MIDI message example:<br/>
  ![image](https://user-images.githubusercontent.com/62774242/180223617-8b22f9c2-8b2c-45d7-9475-a18b21ab67dc.png)

2. The sound of the received note from User A plays on User B's computer.

For more info about the structure of this project and relaled tools, check out these folders: [diagrams](https://github.com/rog-golang-buddies/realtime-midi/tree/main/docs/diagrams), [research](https://github.com/rog-golang-buddies/realtime-midi/tree/main/docs/research)

## What is MIDI?
> MIDI (/ˈmɪdi/; Musical Instrument Digital Interface) is a technical standard that describes a communications protocol, digital interface, and electrical connectors that connect a wide variety of electronic musical instruments, computers, and related audio devices for playing, editing, and recording music.  The specification originates in the paper Universal Synthesizer Interface published by Dave Smith and Chet Wood of Sequential Circuits at the 1981 Audio Engineering Society conference in New York City. A MIDI recording is not an audio signal, as with a sound recording made with a microphone. It is more like a piano roll, indicating the pitch, start time, stop time and other properties of each individual note, rather than the resulting sound.

TL;DR
MIDI is a standard format for sending music notes to other devices or instruments.

## What about the instruments without MIDI support?
We've got you covered 🥳🎉<br/>
What if we convert the sound from your instrument into MIDI and then send it to other users? 😉<br/>
Here is an example which generates MIDI numbers of the notes from the audio received from your microphone.<br/>
[examples/audio-to-midi](examples/audio-to-midi)

# Examples!

- [examples/p2p-webrtc](examples/p2p-webrtc/)
- [examples/time-sync](examples/time-sync/)
- [examples/wspingpong](examples/wspingpong/)
- [examples/ws-noughts-crosses](examples/ws-noughts-crosses)
- [examples/audio-to-midi](examples/audio-to-midi)
