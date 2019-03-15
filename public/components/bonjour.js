import { PolymerElement } from 'https://unpkg.com/@polymer/polymer/polymer-element.js?module'
import { css, dom, doc } from '../utils/template.js'
import './cheer-bg.js'

const style = css`
    :host {
        display: block;
        padding: 1em;
    }
    ::slotted([slot=title]), ::slotted([slot=sub-title]) {
        text-align: center;
    }
`

const body = dom`
    <i-cheer-bg maximum=1>
        <slot name="title"></slot>
    </i-cheer-bg>
    <slot></slot>
`

/**
* @polymer
* @extends HTMLElement
*/
class Bonjour extends PolymerElement {
    static get template () {
        return doc({ style, body })
    }
}

customElements.define('i-bonjour', Bonjour)


