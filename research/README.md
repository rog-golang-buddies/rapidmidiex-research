# SIP

Session Initiation Protocol

- SIP is typically used for VoIP: SIP-phones
- SIP is a signaling protocol
- SIP has many similarities with HTTP and SMTP, like
    - it's text-based
    - **Requests** have methods (`REGISTER`, `INVITE`, `ACK`, `BYE`, ...) similar to *HTTP-methods* (`GET`, `POST`, ...)
    - **Responses** have **status codes** (`1xx`, `200`: success, `3xx`: redirection, `5xx`: server errors, ...)
- SIP can use UDP, TCP or SCTP as transport protocols
    - TCP/UDP port numbers:
        - 5060: non-encrypted
        - 5061: encrypted with TLS
- SIP UA's: Each UA (User Agent) is both client and server
    - UAC: User Agent Client
    - UAS: User Agent Server
    - UAC- and UAS-roles only last for duration of a SIP transaction
- SIP Registrar: location service

Sources:

- https://en.wikipedia.org/wiki/Session_Initiation_Protocol

# SDP 

Session Description Protocol

- SDP is used as the payload of some SIP-messages
- SDP can be used by 2 endpoints to negotiate network metrics, media types, ... (the *session profile*)

# RTP

= Real-time Transport Protocol

> Also **SRTP**: Secure RTP (with **TLS**)

- for transmission of multimedia formats
- has different *profiles* or *payload formats*, like
    - for Audio and video conferences
    - for `H.265`-encoded video
    - **!!! for MIDI !!!**
    - ...

Golang:

- https://libs.garden/go/search?q=rtp

# RTCP 

= RTP Control Protocol

- *helps* RTP-sessions (provides out-of-band statistics and control information)
- mainly for **QoS** (Quality of Service)
- statistics:
    - packet counts
    - packet loss
    - packet delay variation (**jitter**)
    - round-trip delay time (**RTT** / **ping-time**)

sources:

- https://en.wikipedia.org/wiki/RTP_Control_Protocol
- https://developer.mozilla.org/en-US/docs/Web/API/WebRTC_API/Intro_to_RTP

# RTP-MIDI

- how to send MIDI over RTP

- https://en.wikipedia.org/wiki/RTP-MIDI
    - AKA AppleMIDI
    - > Compared to MIDI 1.0, RTP-MIDI includes new features like session management, device synchronization and detection of lost packets, with automatic regeneration of lost data.
    - inside macOS since 2005
    - https://developer.apple.com/library/archive/documentation/Audio/Conceptual/MIDINetworkDriverProtocol/MIDI/MIDI.html
- https://www.midi.org/specifications/midi-transports-specifications/rtp-midi
- https://john-lazzaro.github.io/rtpmidi/
    - a paper from 2004 by AES (Audio Engineering Society) about a very similar use-case: https://john-lazzaro.github.io/sa/pubs/pdf/aes117.pdf
- https://www.rfc-editor.org/rfc/rfc4696.txt
    - *An Implementation Guide for RTP MIDI*
    - November 2006
    - talks about only 2 participants without firewalls and NAT --> not our use-case
- https://www.rfc-editor.org/rfc/rfc6295.txt
    - *RTP Payload Format for MIDI*
    - June 2011
    - > interoperable MIDI networking might foster network music performance applications, in which a group of musicians located at different physical locations interact over a network to perform as  they would if they were located in the same room
    - > these applications have not yet reached the mainstream.  However, experiments in academia and industry continue

Golang:

- https://github.com/laenzlinger/go-midi-rtp

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
