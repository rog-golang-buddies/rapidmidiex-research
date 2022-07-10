import Game from "./components/Game.js";
import { html, render, tw } from "./imports.js";

function App({ name }) {
    return html`
        <main class="${tw`h-screen bg-purple-400 flex items-center justify-center`}">
            <${Game} />
        </main>
    `;
}

render(html`<${App} name="World" />`, document.body);

// need to see why `esm.sh` didn't work

// https://github.com/tw-in-js/twind

// https://esm.sh/?test