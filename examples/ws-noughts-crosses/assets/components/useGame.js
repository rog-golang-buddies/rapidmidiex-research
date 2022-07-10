import { useState } from "../imports.js";

const defaultGameState = () => ({ board: [0, 0, 0, 0, 0, 0, 0, 0, 0], turn: 0, winner: 0, moves: 9 });

/**
 * @param {object} params 
 * @param {(0|1|2)[]} params.board 
 * @param {(1|0)} params.turn 
 * @param {(0|1|2)} params.winner 
 * @returns
 */
export function useGame({ board, turn, winner, ws }) {
    /** @type {[{ board:(0|1|2)[]; winner:0|1|2; moves:number; turn:0|1 }, function]} */const [{ board: boardState, winner: winnerState, moves, turn: turnState }, setGrid] = useState(defaultGameState);

    const onMessage = (/**@type {MessageEvent<any>}*/e) => {
        const message = JSON.parse(e.data) || {};
        switch (message["type"]) {
            case "join":
                console.log(message["data"]);
                break;
            case "reset":
                setGrid(defaultGameState);
                break;
            case "play":
                let { turn: msgTurn, pos: msgPos, moves: msgMoves } = message["data"];
                setGrid(function (/** @type {{ board:(0|1|2)[]; }} */{ board: oldBoard }) {
                    oldBoard[msgPos] = [1, 2][msgTurn ^= 1];
                    return { board: [...oldBoard], turn: msgTurn, winner: hasWinner(oldBoard), moves: --msgMoves };
                });
                break;
            default:
                console.error("unknown message: ", message);
        }
    };

    /** 
     * @param {Event} e
     * @returns {void} 
     */
    const onPlay = (e) => {
        const message = {
            type: "play",
            data: {
                turn: turnState,
                pos: [...e.target.parentNode.children].indexOf(e.target),
                moves,
            }
        };

        ws.send(JSON.stringify(message)); // to server

        // TODO - can only make move if its their turn
        setGrid(function (/** @type {{ board:(0|1|2)[]; turn:1|2; moves:number; }} */{ board: oldBoard, turn: oldTurn, moves }) {
            oldBoard[message.data.pos] = [1, 2][oldTurn ^= 1];
            return { board: [...oldBoard], turn: oldTurn, winner: hasWinner(oldBoard), moves: --moves };
        });
    };

    const onReset = _ => {
        const message = { type: "reset", data: null };
        ws.send(JSON.stringify(message));

        setGrid(defaultGameState);
    };

    return { board: boardState, winner: winnerState, moves, onPlay, onReset, onMessage };
}

/**
 * Algorithm to determine if there is a winner yet or not
 * 
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