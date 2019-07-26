# Down

Down checks if your internet connection is down.

I made this so I no longer have to type out `ping google.com`, wait a bit, and
then mentally parse if it's working or not. It is quicker and more obvious. It
does it with pretty colours. Green good, red bad. Green "UP", red "DOWN".

It is very simple and *very* non-exhaustive. It simply makes a GET request to
Google, Facebook, Amazon and a few more. If any request fails, it will say that
your internet is down. If any succeed, it will say that your internet is up. Or
it will time out after two seconds.

## Installation

Run `go get -u github.com/patrickmcnamara/down`.

Or download a binary executable from the GitHub releases page.

## Usage

Run `down` after installation.
