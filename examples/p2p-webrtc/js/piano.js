// https://jazz-soft.net/demo/Knobs.html
JZZ.synth.Tiny.register();

JZZ.input
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
  .connect(JZZ.input.Kbd({ at: "piano" }).connect(JZZ().openMidiOut()))
  .connect(print);

const midiMsgArea = document.getElementById("midi-msg-area");

const text = [];

function print(msg) {
  text.push(JZZ.MIDI(msg).toString());
  if (text.length > 20) text = text.slice(1);
  midiMsgArea.innerHTML = text.join("<br>");
  midiMsgArea.scrollTop = midiMsgArea.scrollHeight;
  console.log({ msg });
  peerConnection.send(msg);
}
