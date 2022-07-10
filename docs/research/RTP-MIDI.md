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
