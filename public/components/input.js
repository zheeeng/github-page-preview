import { PolymerElement, html } from 'https://unpkg.com/@polymer/polymer/polymer-element.js?module'

/**
* @polymer
* @extends HTMLElement
*/
class Input extends PolymerElement {

    static get properties () {
        return {
          // Configure owner property
          owner: {
            type: String,
              value: 'Daniel',
              reflectToAttribute: true,
          }
        };
      }

    static get template() {
        return html`
        <style>
          div {
            display: inline-block;
            background-color: #ccc;
            border-radius: 8px;
            padding: 4px;
          }
        </style>
        <form on-submit="alert">
            <input type="text" value="{{owner}}" on-input="updateText" />
            <input type="submit" />
        </form>
        <br />
        <div>
          <slot></slot>
        </div>
        <br />
        <br />
        <div>
          [[owner]] send you: {{greeting}}
        </div>
        <!-- bind to the "owner" property -->
        `;
    }

    updateText (event) {
      this.greeting = event.target.value
    }

    alert (event) {
      event.preventDefault()
      alert(this.greeting)
    }

    constructor() {
        super()
        this.greeting = 'Hello, I\'m ' + this.owner
        this.textContent = 'Greeting board:'
    }
}
// Register the new element with the browser
customElements.define('i-input', Input)
