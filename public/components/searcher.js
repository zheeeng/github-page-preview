import { PolymerElement } from 'https://unpkg.com/@polymer/polymer/polymer-element.js?module'
import { css, dom, doc } from '../utils/template.js'

const style = css`
    form {
        max-width: 420px;
        width: 100%;
        margin: 1em auto;
    }
    .form-item {
        position: relative;
        bottom: 0;
        left: 0;
        margin: 2em 0;
    }
    label, input {
        padding-bottom: .5em;
    }
    label {
        display: block;
        position: absolute;
    }
    input {
        display: block;
        width: 100%;
        background: transparent;
        border: 0px solid transparent;
        border-bottom: 1px solid rgba(255,255,255,0.25);
        transition: border-bottom-color 0.15s ease-in-out;
        color: #fff;
        font-size: 1em;
        outline: none;
    }
    input:hover, input:focus {
        border-bottom-color: rgba(255,255,255,0.5);
    }
    input+label {
        bottom: 0;
        left: 0;
        opacity: 0.5;
        pointer-events: none;
        transition: all 0.15s ease-in-out;
    }
    input:hover+label {
        opacity: 0.8;
    }
    input:focus+label, input[data-has-text]+label {
        opacity: 0.8;
        bottom: 1.5em;
    }
`

const body = dom`
    <form>
        <div class="form-item">
            <input
                id="repo-input"
                type="search"
                role="search"
                title="search github repo"
                aria-label="search github repo"
                autocomplete="off"
                value="{{searchRepo}}"
                on-input="handleInput"
                data-has-text$="{{hasText(searchRepo)}}"
            >
            <label for="repo-input">Repo URL (with http:// or https://):</label>
        </div>
        <div class="form-item">
            <input
                id="repo-input"
                type="search"
                role="search"
                title="search github repo"
                aria-label="search github repo"
                autocomplete="off"
                value="{{searchFolder}}"
                on-input="handleInput2"
                data-has-text$="{{hasText(searchFolder)}}"
            >
            <label for="repo-input">For hosting folder</label>
        </div>
    </form>
`

/**
* @polymer
* @extends HTMLElement
*/
class Searcher extends PolymerElement {
    static get template () {
        return doc({ style, body })
    }

    constructor () {
        super()
        this.searchRepo = ''
        this.searchFolder = ''
    }

    hasText (text) {
        return text != ''
    }

    handleInput (event) {
        this.searchRepo = event.target.value
    }
    handleInput2 (event) {
        this.searchFolder = event.target.value
    }
}

customElements.define('i-searcher', Searcher)


