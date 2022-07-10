# Introduction

In this example we try an alternative websocket-implementation: https://github.com/nhooyr/websocket (check there to see more differences with gorilla's and others)

For this example, we simply want to measure latency and it seems this library already provides it:

From the docs at https://pkg.go.dev/nhooyr.io/websocket#Conn.Ping:

> Ping sends a ping to the peer and waits for a pong. Use this to measure latency or ensure the peer is responsive. Ping must be called concurrently with Reader as it does not read from the connection but instead waits for a Reader call to read the pong.

> TCP Keepalives should suffice for most use cases.

# TODO

[x] A simple custom http-server with just one route
[ ] Upgrade client-connections to websocket-connections
[ ] Store all connected clients and show them all on page refresh
[ ] Add a bubbletea-TUI
[ ] ...

# Specs

## no mux

It might be a bit contrarian but this example will only be using **one route** and therefore we do not need a **mux**. 

## custom server

We create a custom `net/http`-server in stead of using the default `http.ListenAndServe(...)` because (according to (source: Jon Bodner - Learning Go, page 251):

- other libraries might use these `http`-package-level-variables already and interfere
- we cannot set our own timeouts

## block favicon

It just pollutes our logging?




