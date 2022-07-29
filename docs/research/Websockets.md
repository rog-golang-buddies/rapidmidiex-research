# Websockets

## TLDR

Websockets can be called a *wire*-protocol where the smallest piece of communication is called the **frame**.
For security, **all frames from client to server** are **masked** (i.e. XOR-encrypted).
(Server-to-client frames are NOT masked.)
The key used for *masking* is chosen by the client **for each frame**.

**Data-frames** can be **Text**-frames or **Binary**-frames.
An application-specific websocket-**message** could be sent over multiple **data-frame**'s (especially bigger messages).

**Control-frames** (like **Ping-**, **Pong** and **Close**-frames) are always <= 125 bytes.

Upon handshake (which happens over the `http(s)`-protocol), the client can negotiate **subprotocols** (like `wamp`) and **extensions**. 
Of course, an application developer can choose their own protocol in stead of one of the existing ones.


## Protocol

https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API/Writing_WebSocket_servers

  - this is basically a synopsis of the **RFC6455**-spec, talks about
  - The Websocket-handshake
    - client handshake requests
    - server handshake responses
    - keeping track of clients
    - `258EAFA5-E914-47DA-95CA-C5AB0DC85B11` is a special UUID used by websocket-servers to prove to the client a handshake was received
  - Data frames
    - *frames* are in fact just a bunch of bytes that we could call a *packet* or a *message* but depending on context one uses the *frame*-moniker
      - Similar to how one speaks about *Ethernet-frames* (layer 2 *data* that travel over a *wire*) but *IP-packets* (layer 3 *bytes* that travel through IP-routers)
    - big websocket **message**s can be split (fragmented) over several **frame**s.
      - We could also say a *frame* is a *wire*-format.
  - Pings and pongs (heartbeat)
  - Extensions (to the base websocket-protocol)
  - Subprotocols (what structure the ws-messages have, compare with XML schema)


Frame format:

```
      0                   1                   2                   3
      0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
     +-+-+-+-+-------+-+-------------+-------------------------------+
     |F|R|R|R| opcode|M| Payload len |    Extended payload length    |
     |I|S|S|S|  (4)  |A|     (7)     |             (16/64)           |
     |N|V|V|V|       |S|             |   (if payload len==126/127)   |
     | |1|2|3|       |K|             |                               |
     +-+-+-+-+-------+-+-------------+ - - - - - - - - - - - - - - - +
     |     Extended payload length continued, if payload len == 127  |
     + - - - - - - - - - - - - - - - +-------------------------------+
     |                               |Masking-key, if MASK set to 1  |
     +-------------------------------+-------------------------------+
     | Masking-key (continued)       |          Payload Data         |
     +-------------------------------- - - - - - - - - - - - - - - - +
     :                     Payload Data continued ...                :
     + - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - +
     |                     Payload Data continued ...                |
     +---------------------------------------------------------------+
```
- `FIN` (1 bit): set when this **frame** is the end of a **message**, if `0` the server keeps listening for more parts of the message
- `RSV1/2/3`: only used by extensions
- `MASK` (1 bit): set to indicate the messages sent by the client (**client -> server**) are *masked* AKA *XOR-encrypted*
  - all messages from the client should be masked!
    - client should discard received masked messages
    - server should discard received unmasked messages and immediately close connection
  - client needs to set the `Masking-key` with which it **encoded** the payload
  - server will need to read the `Masking-key` to **decode** the payload
- `opcode` (4 bits): indicates what kind of frame this is
  - `0x0`: continuation frame
  - `0x1`: text frame
  - `0x2`: binary frame
  - `0x8`: connection close frame
  - `0x9`: ping frame
  - `0xA`: pong frame
- `payload len` (7 bits): enough for when payload `< 125`
  - `126`: more payload length-bits to fetch from `extended payload length`
  - `127`: more payload length-bits to fetch from `extended payload length continued`


Example of message-flow (1 message sent over 4 frames):

```
Client: FIN=1, opcode=0x1, msg="hello"
Server: (process complete message immediately) Hi.
Client: FIN=0, opcode=0x1, msg="and a"
Server: (listening, new message containing text started)
Client: FIN=0, opcode=0x0, msg="happy new"
Server: (listening, payload concatenated to previous message)
Client: FIN=1, opcode=0x0, msg="year!"
Server: (process complete message) Happy new year to you too!
```
- only text- (opcode `0x1`) and binary-frames (opcode `0x2`) can be fragmented
- opcode `0x0` means this payload should be added to the previous one to complete the message
- opcode `0x0` with `FIN`=`1` means the message is complete and can be processed
  - apparently a complete message can be added to anyway?
- `pings` and `pongs`
  - they are **control frames**
  - when receiving a ping, other side should send a pong ASAP (if connection still open)
  - pong uses same *payload* as what was sent by ping
  - max. payload: `125`
- closing handshake
  - peer sends a `0x8`-control frame
  - other peer replies
  - first peer closes connection


## Some notes from **RFC6455**:

- The Handshake
  - The Request
    - required headers:
      - `Request-URI`
      - `Host` indicates servers authority (important for Same-Origin Policy / cross-site scripting) (RFC6454)
      - `Origin` (only required if coming from a **browser** client)
      - `Connection: Upgrade`
      - `Sec-WebSocket-Key: ...` (base64-encoded 16-byte value)
      - `Sec-WebSocket-Version: 13`
    - optional headers:
      - `Sec-WebSocket-Protocol`: comma-separated list of *Subprotocols* (they **structure** the **websocket-payload** (e.g. `wamp`)) the client wishes to speak
        - the server can only reply with **1 subprotocol** it will speak with the client
      - `Sec-WebSocket-Extensions`: which protocol-level extensions (that **modify** the **websocket-payload**) the client wishes to speak
      - cookies
  - The Response
    - required headers:
      - `101`: switching protocols
      - `Upgrade: websocket`
      - `Connection: Upgrade`
      - `Sec-WebSocket-Accept`: a computed field:
        - concat key with `Sec-WebSocket-Key` from client with the fixed value `258EAFA5-E914-47DA-95CA-C5AB0DC85B11`
        - take *SHA-1*-hash
        - *base64*-encode
    - optional headers:
      - `Sec-WebSocket-Protocol`
      - `Sec-WebSocket-Extensions` (multiple extensions can be used)
- **Masking**
  - Why?
    - security for attacks on *infrastructure*: proxies could alter client-messages, send fake messages, poison caches, ...
  - How?
    - client choses a random key **for each frame**
      - key must be cryptographically secure AKA from strong source of entropy (RFC4086)
        - i.e. not predictable from masking keys used in previous frames
      - masking-algorithm is XOR-based, doesn't change payload-length

## More information

- https://lucumr.pocoo.org/2012/9/24/websockets-101/
  - by **Armin Ronacher**, creator of flask
  - Keep in mind: it's an OLD document (**2012**)
  - advises to always use **websocket through TLS** so intermediates can't mess it up
  - different URL-scheme ( `ws(s)`) with specific grammar
    - no anchors (`#foo`)
  - problems with HTTP/TCP-proxies
  - doesn't like that ping/pong has payload (the websocket-browser-API doesn't add payload?)
  - closing handshake: client is supposed to give server some time to close
    - Firefox cares and will also reconnect after TCP disonnection?
  - if no browser required, just speak TCP?
- https://hpbn.co/websocket/
  - one chapter of the book **High Performance Browser Networking** by **Ilya Grigorik** (publisher: O'Reilly)
  - https://www.igvita.com/
- the book **WebSocket, Lightweight Client-Server Communications** by **Andrew Lombardi** (publisher: O'Reilly)


## Reference

https://www.iana.org/assignments/websocket/websocket.xml

Check this for:

- supported subprotocols like `wamp`
- version numbers (we're currently at **13** with **RFC6455**)
- close codes (`1000`: normal closure, `1002`: protocol error, ...)
- opcodes (`TextFrame`, `BinaryFrame`, `PingFrame`, `PongFrame`, ...)
