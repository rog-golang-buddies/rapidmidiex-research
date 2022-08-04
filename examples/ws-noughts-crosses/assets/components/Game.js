import { html, useState, tw, useEffect } from "../imports.js";
import { useGame } from "./useGame.js";


export default function Game() {
    const { board, winner, moves, onPlay, onReset } = useGame({ board: [0, 0, 0, 0, 0, 0, 0, 0, 0], turn: 1, winner: 0, url: "ws://localhost:8888/play" });

    // useEffect(() => {
    //     ws.addEventListener('open', e => console.log("connection opened"));
    //     ws.addEventListener('close', e => console.log("closed"));
    //     ws.addEventListener("error", e => console.error(e));
    //     ws.addEventListener('message', onMessage);
    // }, []); //have no idea why this needs to be here

    return html`
    <div class=${tw`flex-col`}>
        <div class=${tw`grid grid-cols-3 gap-3`} onClick=${onPlay}>
            ${board.map(cell => html`<button disabled=${winner || cell} class=${tw`border-green-200 border-2 p-2`}>${cell}</button>`)}
        </div>

        <p>Game status: ${winner}</p>
        <p>Moves left: ${moves}</p>
        
        <button onClick=${onReset} class=${tw`p-2 bg-gray-200`} disabled=${!winner && moves}>Reset Game</button>
    </div>
    `;
}
