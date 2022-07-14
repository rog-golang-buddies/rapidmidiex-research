# Introduction

Here we collect (and learn about) all possible useful information.

When programming Internet-applications, some basic networking knowledge about https://en.wikipedia.org/wiki/Internet_protocol_suite can be useful.

When it comes to streaming audio/video or other realtime communications, some specific protocols have been developed over the years. We explore them a bit in [StreamingProtocols](StreamingProtocols.md).

One particular interesting use-case that has been researched by audio professionals (and is also implemented by Apple in their OS's and therefore also known as **AppleMIDI**) is [RTP-MIDI](RTP-MIDI.md).

All previous links serve more as background-information.

Next we collect what could actually be useful for our project:

# WebRTC

- uses **SCTP**: Stream Control Transmission Protocol (transport layer protocol)
- uses **DTLS**: TLS for UDP- or SCTP-datagrams
- several Javascript API's:
  - `getUserMedia`
  - `RTCPeerConnection`
  - `RTCDataChannel` (similar API to websocket but latency)
  - `getStats`
- requires SIP, **SIP over websockets (RFC7118)**, XMPP (Jabber), ... for signaling

Golang:

- https://libs.garden/go/search?q=webrtc

# MIDI

Golang:

- https://libs.garden/go/search?q=midi
- https://github.com/gomidi/midi
  - contains a **WASM**-example that can talk to real midi-hardware from the webbrowser!

# Music theory / math

Golang:

- https://github.com/go-music-theory/music-theory

# Front-end audio/midi

- https://jazz-soft.net/doc/
  - used in [p2p-webrtc example](../examples/p2p-webrtc/index.html)

- https://github.com/joshreiss/Working-with-the-Web-Audio-API
  - many examples in tutorial-format about WebAudio
  - chapter 18 and 19 are about `AudioWorklet`'s (https://developer.mozilla.org/en-US/docs/Web/API/Web_Audio_API/Using_AudioWorklet)

- https://web.dev/audio-scheduling/
  - don't use `setTimeout` but `requestAnimationFrame` or *custom system?* 


# CLI audio processing utility

- https://github.com/faiface/beep
