# Lemmy Wails Demo

This is a `wails` project using [Lemmy Wails](https://github.com/LouisBrunner/lemmy-wails).

## Installation

Use `make install` to prepare this project.

## Live Development

To run in live development mode, run `make dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `make build`.

## Linting & Testing

Run `make lint` or `make test`
