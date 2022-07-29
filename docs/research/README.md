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
  - https://pion.ly/

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

# Websockets

See [Websockets.md](Websockets.md) for some notes on the Websocket-protocol as per RFC's and other useful websites.

Golang:

- https://pkg.go.dev/golang.org/x/net/websocket
  - (per documentation:) *lacks features*
- https://github.com/gorilla/websocket
  - (per documentation:) *complete and tested*, passes *Autobahn Test Suite*
    - https://github.com/crossbario/autobahn-testsuite
  - probably the most used websocket-implementation
- https://github.com/nhooyr/websocket
  - passes *Autobahn Test Suite*
  - Some interesting features that **might** make it more useful for our project
    - Zero alloc reads and writes
    - Supports compiling to **WASM**
    - solves a problem with closing handshake that gorilla still has (https://github.com/gorilla/websocket/issues/448)
- https://github.com/gobwas/ws
  - written for *mail.ru* to handle millions of users checking their e-mail, see https://www.freecodecamp.org/news/million-websockets-and-go-cc58418460bb
  - speaks of *Zero Copy HTTP Upgrades* (https://github.com/gobwas/ws#zero-copy-upgrade)
    - this is when **one** client's HTTP-connection gets upgraded to a `ws`-connection
    - important for when you have many millions of clients checking e-mail
    - not very important for our use case since a jamming session would reasonably have only about a dozen people
      - unless we want to make it **MMOJ** (Massive Multiplayer Online Jamming) :wink: but even then the *joining a jam-session*-part is not the real bottleneck

# Websocket performance

- https://centrifugal.github.io/centrifugo/blog/scaling_websocket/
  - setting up web sockets for scalability
  - basically this article is promotion for https://github.com/centrifugal/centrifuge
  - talks about
    - the different golang-libraries, author prefers gorilla
    - OS-tuning: a file descriptor per connection? 
      - https://docs.riak.com/riak/kv/2.2.3/using/performance/open-files-limit.1.html
      - linux: `ulimit -n`
      - macos: `launchctl limit maxfiles`
      - or: https://pkg.go.dev/golang.org/x/net/netutil#LimitListener
      
      - or some low-level TCP/IP-settings: https://gist.github.com/mustafaturan/47268d8ad6d56cadda357e4c438f51ca
    - pub/sub-brokers (RabbitMQ, Kafka, Redis, ...)
    - massive reconnect problem
    - benefits of using **message event stream** (a buffer with client-state)
  - https://crossbario.com/blog/Dissecting-Websocket-Overhead/
    - some benchmarks on a GbE-switch

# Utility

- https://pkg.go.dev/golang.org/x/time/rate
  - used in websocket-chat-example: https://github.com/nhooyr/websocket/blob/master/examples/chat
