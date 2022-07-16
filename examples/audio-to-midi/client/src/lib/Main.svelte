<script lang="ts">
  // audio context to control audio input
  let audioContext: AudioContext = null;

  // variable to store Audio Context Analyser Node
  let analyser: AnalyserNode = null;

  // variable to store microphone stream data
  let mediaStreamSource: MediaStreamAudioSourceNode = null;

  // false if note was ignored
  let isConfident: boolean = false;

  // defines the sensitivity of the algorithm
  // higher value means lower sensitivity
  let sensitivity: number = 0.05; // default is 0.05

  // length of an octave which has 12 notes
  // in Western musical scale
  const octaveLength = 12;

  // pitch is the same as frequency, just different names
  let pitch: number = 0;

  // interface for defining note
  interface Note {
    Name: string;
    Octave: number;
  }

  // the name of the note which is chosen from noteStrings array
  let note: Note = {
    Name: "A",
    Octave: 4,
  };

  // current octave in use
  let octave: number = 4;

  // the amount of frequency between the detected note that is the closest to the frequency
  // imagine note A4. it has the frequency of 440 Hz. if the frequency of the sound is 430 Hz
  // it's still the same note which is A4. but with -10 Hz of deviation.
  let deviation: number = 0;

  // array for storing the previous detected notes
  let noteHistory: Note[] = [];

  // the constant length of noteHistory
  const historyLength = 10;

  // variable for disabling start button when clicked
  let startButtonDisabled: boolean = false;

  // boolean to control settings window display
  let showSettings: boolean = false;

  // string array of 9 octaves of notes
  const noteStrings = [
    "C",
    "C#",
    "D",
    "D#",
    "E",
    "F",
    "F#",
    "G",
    "G#",
    "A",
    "A#",
    "B",
  ];

  // brain of the tuner
  // don't change anything if you don't know how the physics of signals work
  function autoCorrelate(buf, sampleRate) {
    // Implements the ACF2+ algorithm
    let SIZE = buf.length;
    let rms = 0;

    for (let i = 0; i < SIZE; i++) {
      let val = buf[i];
      rms += val * val;
    }
    rms = Math.sqrt(rms / SIZE);
    if (rms < sensitivity)
      // not enough signal
      // the note is ignored
      return -1;

    let r1 = 0,
      r2 = SIZE - 1,
      thres = 0.2;
    for (let i = 0; i < SIZE / 2; i++)
      if (Math.abs(buf[i]) < thres) {
        r1 = i;
        break;
      }
    for (let i = 1; i < SIZE / 2; i++)
      if (Math.abs(buf[SIZE - i]) < thres) {
        r2 = SIZE - i;
        break;
      }

    buf = buf.slice(r1, r2);
    SIZE = buf.length;

    let c = new Array(SIZE).fill(0);
    for (let i = 0; i < SIZE; i++)
      for (let j = 0; j < SIZE - i; j++) c[i] = c[i] + buf[j] * buf[j + i];

    let d = 0;
    while (c[d] > c[d + 1]) d++;
    let maxval = -1,
      maxpos = -1;
    for (let i = d; i < SIZE; i++) {
      if (c[i] > maxval) {
        maxval = c[i];
        maxpos = i;
      }
    }
    let T0 = maxpos;

    let x1 = c[T0 - 1],
      x2 = c[T0],
      x3 = c[T0 + 1];
    let a = (x1 + x3 - 2 * x2) / 2,
      b = (x3 - x1) / 2;
    if (a) T0 = T0 - b / (2 * a);

    return sampleRate / T0;
  }

  // an async function that waits for the user to grant microphone permission
  async function getUserMedia() {
    navigator.mediaDevices
      .getUserMedia({
        audio: {
          echoCancellation: false,
          autoGainControl: false,
          noiseSuppression: false,
        },
        video: false,
      })
      .then((stream) => {
        // run the necessary commands when permission was granted
        gotStream(stream);
        startButtonDisabled = true;
      })
      .catch((err) => {
        // display error if permission was not granted
        alert("getUserMedia threw exception:" + err);
        startButtonDisabled = false;
      });
  }

  function gotStream(stream) {
    // Create an AudioNode from the stream.
    // this is the stream of sound received from microphone
    mediaStreamSource = audioContext.createMediaStreamSource(stream);

    // Connect it to the destination.
    // this is like the tools needed for analaysing the sound buffer
    // more info here: https://developer.mozilla.org/en-US/docs/Web/API/BaseAudioContext/createAnalyser
    analyser = audioContext.createAnalyser();
    analyser.fftSize = 2048;

    // connect the analyser to audio stream
    mediaStreamSource.connect(analyser);

    // start detecting notes
    updatePitch();
  }

  // converts frequency to note
  // frequency of 440 will be converted to note 'A'
  // more info: https://alijamieson.co.uk/2021/12/20/describing-relationship-two-notes/#:~:text=An%20octave%20is%20an%20intervals,A5%20would%20be%20880%20Hz.
  function noteFromPitch(frequency) {
    var noteNum = octaveLength * (Math.log(frequency / 440) / Math.log(2));
    return Math.round(noteNum) + 69;
  }

  // note: the number 69 corresponds to the pitch A4
  // more info: https://www.audiolabs-erlangen.de/resources/MIR/FMP/C1/C1S3_FrequencyPitch.html
  function frequencyFromNoteNumber(note) {
    return 440 * Math.pow(2, (note - 69) / octaveLength);
  }

  // an octave has 12 notes and 1200 cents
  // which means that there is 100 cents between each note
  // cents off from pitch gives us the deviation from the detected note
  // if it's higher than 50 or lower than -50 it means we have entered the bounds of the other notes
  // eg: out of tune
  function centsOffFromPitch(frequency, note) {
    return Math.floor(
      (octaveLength *
        100 *
        Math.log(frequency / frequencyFromNoteNumber(note))) /
        Math.log(2)
    );
  }

  // array for received buffer of audio
  let buflen = 2048;
  let buf = new Float32Array(buflen);

  // updates the note using requestAnimationFrame
  function updatePitch() {
    analyser.getFloatTimeDomainData(buf);
    let ac = autoCorrelate(buf, audioContext.sampleRate);

    if (ac == -1) {
      // note was ignored
      isConfident = false;
    } else {
      isConfident = true;
      pitch = ac;

      // the index of the detected note
      let noteIdx = noteFromPitch(pitch);

      if (note?.Name !== noteHistory[noteHistory.length - 1]?.Name) {
        // keep the length of the array fixed
        if (noteHistory.length === historyLength) {
          // remove first note of noteHistory array and shift other notes one index to the left
          noteHistory.shift();
        }

        ws.send(
          JSON.stringify({
            midi_num: noteIdx,
            state: true,
          })
        );

        // new note detected, push it to history array
        noteHistory = [
          ...noteHistory,
          { Name: note?.Name, Octave: note.Octave },
        ];
      }

      // noteIdx % noteString.length(108) is one octave high (because octaves start from 0)
      // -12 decreases the octave
      note.Name = noteStrings[noteIdx % noteStrings.length];
      note.Octave = Math.floor(noteIdx / octaveLength) - 1;
      octave = note.Octave;
      deviation = centsOffFromPitch(pitch, noteIdx); // deviation from original note frequency
    }
    requestAnimationFrame(updatePitch);
  }

  let started: boolean = false;
  let ws; // variable to store web socket connection
  let connected: boolean = false;

  // establich web socket connnection
  function connectWebSocket() {
    ws = new WebSocket("ws://localhost:8080/midi/play");

    ws.addEventListener("open", () => {
      connected = true;
    });

    ws.addEventListener("close", () => {
      connected = false;
    });
  }

  function init() {
    if (!connected) {
      connectWebSocket();
    }
    started = true;
    audioContext = new (window.AudioContext || globalThis.webkitAudioContext)();
    getUserMedia(); // get microphone permission
  }
</script>

<main>
  <h2>
    {#if connected}
      Connected
    {:else}
      Disconnected
    {/if}
  </h2>
  Connection will close after 30 seconds of inactivity. you need to reconnect before
  you continue.<br />
  <button on:click={connectWebSocket} disabled={connected}>Connect</button>
  <button on:click={init} disabled={started}>Start</button>
  <p>Note: {note.Name}</p>
  <p>Octave: {note.Octave}</p>
  <p>Frequency: {pitch}</p>
  <p>Cents off: {deviation}</p>
</main>

<style lang="scss">
</style>
