import { html, useState, tw } from "../imports.js";


/**
 * @param {object} params 
 * @param {(0|1|2)[]} params.board 
 * @param {(1|0)} params.turn 
 * @param {(0|1|2)} params.winner 
 * @returns
 */
function useGame({ board, turn, winner }) {
    /** @type {[{ board:(0|1|2)[]; winner:0|1|2 }, function]} */const [{ board: aBoard, winner: aWinner, moves, turn: aTurn }, setGrid] = useState({ board, turn, winner, moves: board.length });

    const ws = new WebSocket("ws://localhost:8080/ws");
    ws.addEventListener('open', e => {
        console.log("connection opened");
    });

    ws.addEventListener('message', e => {
        const message = JSON.parse(e.data);

        switch (message["type"]) {
            case "join":
                console.log(message["data"]);
                break;
            case "play":
                let { turn: bTurn, pos: bPos, moves: bMoves } = message["data"];
                setGrid(function (/** @type {{ board:(0|1|2)[]; moves:number; }} */{ board: oldBoard }) {
                    oldBoard[bPos] = [1, 2][bTurn ^= 1];
                    return { board: [...oldBoard], turn: bTurn, winner: hasWinner([...oldBoard]), moves: --bMoves };
                });
                console.log(message["data"]);
                break;
            default:
                console.error("unknown message: ", message);
        }
    });

    ws.addEventListener('close', e => {
        console.log("closed");
    });

    ws.addEventListener("error", e => console.error(e));


    /** 
     * @param {Event} e
     * @returns {void} 
     */
    const play = (e) => {
        console.log(aTurn);
        const message = {
            type: "play",
            data: {
                turn: aTurn,
                pos: [...e.target.parentNode.children].indexOf(e.target),
                moves,
            }
        };

        ws.send(JSON.stringify(message)); // to server

        setGrid(function (/** @type {{ board:(0|1|2)[]; turn:1|2; moves:number; }} */{ board: oldBoard, turn: oldTurn, moves }) {
            // oldBoard[[...e.target.parentNode.children].indexOf(e.target)] = [1, 2][oldTurn ^= 1];
            oldBoard[message.data.pos] = [1, 2][oldTurn ^= 1];
            return { board: [...oldBoard], turn: oldTurn, winner: hasWinner([...oldBoard]), moves: --moves };
        });
    };

    const reset = _ => setGrid(_ => ({ board, turn, winner, moves: board.length }));

    return { board: aBoard, winner: aWinner, moves, play, reset };
}

/**
 * @param {number[]} board
 * @returns
 */
const hasWinner = (board) => {
    // very hacky solution
    // a,b,c
    // 0,0,0
    // 0,0,1
    // 0,1,0
    // 0,1,1
    // 1,0,0
    // 1,0,1
    // 1,1,0
    // 1,1,1 -> winner
    // 1^1^1 ->
    const r1 = board[0] === board[1] && board[0] === board[2] ? board[0] : 0;
    const c1 = board[0] === board[3] && board[0] === board[6] ? board[0] : 0;
    const d1 = board[0] === board[4] && board[0] === board[8] ? board[0] : 0;
    const c2 = board[1] === board[4] && board[1] === board[7] ? board[1] : 0;
    const d2 = board[2] === board[4] && board[2] === board[6] ? board[2] : 0;
    const c3 = board[2] === board[5] && board[2] === board[8] ? board[2] : 0;
    const r2 = board[3] === board[4] && board[3] === board[5] ? board[3] : 0;
    const r3 = board[6] === board[7] && board[6] === board[8] ? board[6] : 0;

    return r1 || r2 || r3 || c1 || c2 || c3 || d1 || d2;
};

export default function Game() {
    const { board, winner, moves, play, reset } = useGame({ board: [0, 0, 0, 0, 0, 0, 0, 0, 0], turn: 1, winner: 0 });

    return html`
    <div class=${tw`flex-col`}>
        <div class=${tw`grid grid-cols-3 gap-3`} onClick=${play}>
            ${board.map(cell => html`<button disabled=${winner || cell} class=${tw`border-green-200 border-2 p-2`}>${cell}</button>`)}
        </div>

        <p>Game status: ${winner}</p>
        <p>Moves left: ${moves}</p>
        
        <button onClick=${reset} class=${tw`p-2 bg-gray-200`} disabled=${!winner && moves}>Reset Game</button>
    </div>
    `;
}
