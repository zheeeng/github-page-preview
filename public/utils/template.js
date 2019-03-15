import { html } from 'https://unpkg.com/@polymer/polymer/polymer-element.js?module'

export const css = (...args) => String.raw(...args) ? html`<style>${html(...args)}</style>` : html``

export const dom = html

export const js = (...args) => String.raw(...args) ? html`<script>${html(...args)}</script>` : html``

export const doc = ({ style, body, script }) => html`
  ${style || html``}
  ${body || html``}
  ${script || html``}
`
