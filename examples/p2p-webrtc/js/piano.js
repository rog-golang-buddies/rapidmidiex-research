// https://jazz-soft.net/demo/Knobs.html
JZZ.synth.Tiny.register();

localInstrument = JZZ.input
  .ASCII({
    A: "F#4",
    Z: "G4",
    S: "G#4",
    X: "A4",
    D: "Bb4",
    C: "B4",
    V: "C5",
    G: "C#5",
    B: "D5",
    H: "D#5",
    N: "E5",
    M: "F5",
    K: "F#5",
    "<": "G5",
    L: "G#5",
    ">": "A5",
    ":": "Bb5",
  })
  .connect(
    JZZ.input.Kbd({ at: "piano" }).connect(JZZ().openMidiOut()).connect(print)
  );

const remoteMIDIchannel = 1;
// https://jazz-soft.net/demo/GeneralMidi.html
const remoteMIDIInstrument = 16; // Organ
const remoteInstrument = JZZ.input
  .Kbd({ active: false })
  .connect(JZZ().openMidiOut());
remoteInstrument.program(remoteMIDIchannel, remoteMIDIInstrument);
const midiMsgArea = document.getElementById("midi-msg-area");

function print(msg) {
  const midiString = JZZ.MIDI(msg).toString();
  console.log({ midiMsg: midiString });
  const [channel, note, velocity] = Array.from(msg);
  const midiMsg = {
    channel,
    note,
    velocity,
    isNoteOn: midiString.split("Note")[1].trim() === "On",
  };
  if (peerConnection.dataChannel.readyState === "open") {
    peerConnection.dataChannel.send(JSON.stringify(midiMsg));
  }
}

function onMidiReceived(msg) {
  const { channel, note, velocity, isNoteOn } = JSON.parse(msg.data);
  console.log("Received MIDI message: ", { channel, note, velocity });
  const methodName = isNoteOn ? "noteOn" : "noteOff";

  // Play remote note on 2nd instrument
  remoteInstrument[methodName](remoteMIDIchannel, note, velocity);
}

window.onConnectionOpen = (connection) => {
  connection.dataChannel.onmessage = onMidiReceived;
};
