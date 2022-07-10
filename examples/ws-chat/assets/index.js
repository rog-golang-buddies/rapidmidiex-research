import { html, render, tw } from "./imports.js";

const ws = new WebSocket("ws://localhost:8080/ws");
ws.addEventListener('open', e => {
    console.log("connection opened");
});

ws.addEventListener('message', e => {
    console.log("received from server: ", e.data);

});

ws.addEventListener('close', e => {
    console.log("closed");
});

ws.addEventListener("error", e => console.error(e));

function App({ name }) {


    const onClick = () => {
        const rand = Math.round(Math.random() * 10);
        ws.send(rand);
        // console.log(rand);
    };

    return html`
        <main class="${tw`h-screen bg-purple-400 flex items-center justify-center`}">
            <button onClick=${onClick}>start</button>
        </main>
    `;
}

render(html`<${App} name="World" />`, document.body);

// need to see why `esm.sh` didn't work

// https://github.com/tw-in-js/twind

// https://esm.sh/?test