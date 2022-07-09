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

