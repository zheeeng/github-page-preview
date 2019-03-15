import { PolymerElement } from 'https://unpkg.com/@polymer/polymer/polymer-element.js?module'
import { css, dom, doc } from '../utils/template.js'

const style = css`
    :host {
        display: block;
        width: 200px;
        position: relative;
        overflow: hidden;
    }

    :host([hidden]) {
        display: none !important;
    }

    #progressContainer {
        position: relative;
    }

    #progressContainer,
    /* the stripe for the indeterminate animation*/
    .indeterminate::after {
        height: var(--paper-progress-height, 4px);
    }

    #primaryProgress,
    #secondaryProgress,
    .indeterminate::after {
    }

    #progressContainer,
    .indeterminate::after {
        background: var(--paper-progress-container-color);
    }

    :host(.transiting) #primaryProgress,
    :host(.transiting) #secondaryProgress {
        -webkit-transition-property: -webkit-transform;
        transition-property: transform;
        /* Duration */
        -webkit-transition-duration: var(--paper-progress-transition-duration, 0.08s);
        transition-duration: var(--paper-progress-transition-duration, 0.08s);
        /* Timing function */
        -webkit-transition-timing-function: var(--paper-progress-transition-timing-function, ease);
        transition-timing-function: var(--paper-progress-transition-timing-function, ease);
        /* Delay */
        -webkit-transition-delay: var(--paper-progress-transition-delay, 0s);
        transition-delay: var(--paper-progress-transition-delay, 0s);
    }

    #primaryProgress,
    #secondaryProgress {
        -webkit-transform-origin: left center;
        transform-origin: left center;
        -webkit-transform: scaleX(0);
        transform: scaleX(0);
        will-change: transform;
    }

    #primaryProgress {
        background: var(--paper-progress-active-color);
    }

    #secondaryProgress {
        background: var(--paper-progress-secondary-color);
    }

    :host([disabled]) #primaryProgress {
        background: var(--paper-progress-disabled-active-color);
    }

    :host([disabled]) #secondaryProgress {
        background: var(--paper-progress-disabled-secondary-color);
    }

    :host(:not([disabled])) #primaryProgress.indeterminate {
        -webkit-transform-origin: right center;
        transform-origin: right center;
        -webkit-animation: indeterminate-bar var(--paper-progress-indeterminate-cycle-duration, 2s) linear infinite;
        animation: indeterminate-bar var(--paper-progress-indeterminate-cycle-duration, 2s) linear infinite;
    }

    :host(:not([disabled])) #primaryProgress.indeterminate::after {
        content: "";
        -webkit-transform-origin: center center;
        transform-origin: center center;
        -webkit-animation: indeterminate-splitter var(--paper-progress-indeterminate-cycle-duration, 2s) linear infinite;
        animation: indeterminate-splitter var(--paper-progress-indeterminate-cycle-duration, 2s) linear infinite;
    }

    @-webkit-keyframes indeterminate-bar {
        0% {
            -webkit-transform: scaleX(1) translateX(-100%);
        }

        50% {
            -webkit-transform: scaleX(1) translateX(0%);
        }

        75% {
            -webkit-transform: scaleX(1) translateX(0%);
            -webkit-animation-timing-function: cubic-bezier(.28, .62, .37, .91);
        }

        100% {
            -webkit-transform: scaleX(0) translateX(0%);
        }
    }

    @-webkit-keyframes indeterminate-splitter {
        0% {
            -webkit-transform: scaleX(.75) translateX(-125%);
        }

        30% {
            -webkit-transform: scaleX(.75) translateX(-125%);
            -webkit-animation-timing-function: cubic-bezier(.42, 0, .6, .8);
        }

        90% {
            -webkit-transform: scaleX(.75) translateX(125%);
        }

        100% {
            -webkit-transform: scaleX(.75) translateX(125%);
        }
    }

    @keyframes indeterminate-bar {
        0% {
            transform: scaleX(1) translateX(-100%);
        }

        50% {
            transform: scaleX(1) translateX(0%);
        }

        75% {
            transform: scaleX(1) translateX(0%);
            animation-timing-function: cubic-bezier(.28, .62, .37, .91);
        }

        100% {
            transform: scaleX(0) translateX(0%);
        }
    }

    @keyframes indeterminate-splitter {
        0% {
            transform: scaleX(.75) translateX(-125%);
        }

        30% {
            transform: scaleX(.75) translateX(-125%);
            animation-timing-function: cubic-bezier(.42, 0, .6, .8);
        }

        90% {
            transform: scaleX(.75) translateX(125%);
        }

        100% {
            transform: scaleX(.75) translateX(125%);
        }
    }
`

const body = dom`
    <div id="progressContainer">
        <div id="secondaryProgress" hidden$="[[_hideSecondaryProgress(secondaryRatio)]]"></div>
        <div id="primaryProgress"></div>
    </div>
`

/**
* @polymer
* @extends HTMLElement
*/
class ProgressBar extends PolymerElement {
    static get template () {
        return doc({ style, body })
    }
}

customElements.define('i-progress-bar', ProgressBar)

