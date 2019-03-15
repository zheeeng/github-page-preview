import { PolymerElement, html } from 'https://unpkg.com/@polymer/polymer/polymer-element.js?module'
import 'https://unpkg.com/@polymer/paper-checkbox/paper-checkbox.js?module'

/**
* @polymer
* @extends HTMLElement
*/
class IconToggle extends PolymerElement {
  static get properties() {
    return {
      mystring: {
        type: String,
        value: 'hello world'
      }
    };
  }
  static get template() {
    return html`
      <h1>i am a custom element</h1>
      <p>[[mystring]]</p>
      <paper-checkbox>can has paper-checkbox?</paper-checkbox>
    `;
  }
}

customElements.define('icon-toggle', IconToggle)
