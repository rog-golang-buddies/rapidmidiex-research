import { useState } from "../imports.js";

const defaultGameState = () => ({ board: [0, 0, 0, 0, 0, 0, 0, 0, 0], turn: false, winner: 0, moves: 9 });

/**
 * State logic
 * 
 * @param {object} params -
 * @param {(0|1|2)[]} params.board 
 * @param {boolean} params.turn -- '0' indicates player A & '1' indicates player B, but this could just be a boolean value as it is only
 * @param {0|1|2} params.winner 
 * @param {WebSocket} params.ws 
 * @returns
 */
export function useGame({ board, turn, winner, ws }) {
    /** @type {[{ board:(0|1|2)[]; winner:0|1|2; moves:number; turn:boolean }, function]} */const [{ board: curBoard, winner: curWinner, moves, turn: curTurn }, setGrid] = useState(defaultGameState);

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
                /** @type {{ moves:number; turn:boolean; pos:number }} */let { turn, pos, moves } = message["data"];
                setGrid(function (/** @type {{ board:(0|1|2)[]; }} */{ board }) {
                    board[pos] = (turn = !turn) ? 2 : 1;
                    return { board: [...board], turn, winner: hasWinner(board), moves: --moves };
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
        // don't run if it's not a button
        if (e.composedPath()[0].localName !== "button") return;

        const message = {
            type: "play",
            data: {
                turn: curTurn,
                pos: [...e.target.parentNode.children].indexOf(e.target),
                moves,
            }
        };

        ws.send(JSON.stringify(message)); // to server

        // TODO - can only make move if its their turn
        setGrid(function (/** @type {{ board:(0|1|2)[]; turn:boolean; moves:number; }} */{ board, turn, moves }) {
            board[message.data.pos] = (turn = !turn) ? 2 : 1;
            // board[message.data.pos] = [1, 2][(turn = !turn) | 0];
            return { board: [...board], turn, winner: hasWinner(board), moves: --moves };
        });
    };

    const onReset = _ => {
        const message = { type: "reset", data: null };
        ws.send(JSON.stringify(message));

        setGrid(defaultGameState);
    };

    return { board: curBoard, winner: curWinner, moves, onPlay, onReset, onMessage };
}

/**
 * Algorithm to determine if there is a winner yet or not
 * 
 * @param {number[]} board
 * @returns
 */
const hasWinner = (board) => {
    // very hacky solution
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