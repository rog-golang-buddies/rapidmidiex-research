# Multiplayer Noughts & Crosses

## How to try

- `cd` into project and run `go run ./cmd/`
- open two broswers at `localhost:/8080/` side-by-side
- there are 3 states - **0, 1 & 2** represents empty, player 1 and player 2 respectively
- game buttons disable once a moves has been placed or there is a win/draw state

## Things to consider

There currently is concept of rooms & each player can play both *noughts* and *cross* moves. As a POC this is fine, but would want to handle both these cases.

If a page refreshes, then the old connection disconnects yet the game state doesn't reset. It's worth considering how we would handle what happens when one user looses their connection