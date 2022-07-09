# time-sync-demo

## How to run?

- run _go mod tidy_ to install libraries
- run _go run cmd/main.go_ to start server

## How to use?

- start sync with sending a websocket message to: ws://localhost:8080/time/sync

  message type: JSON, data: `{ "bpm": 120 }`

- stop sync with sending a GET request to: http://localhost:8080/time/stop
