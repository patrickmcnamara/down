# Down

Down checks if your internet connection is down.

I made this so I no longer have to type out `ping google.com`, wait a bit, and
then mentally parse if it's working or not. It is quicker and more obvious. It
does it with pretty colours. Green good, red bad. Green "UP", red "DOWN".

It is very simple and *very* non-exhaustive. It simply makes a GET request to
Google. If that fails, it will say that your internet is down, else it will say
that your internet is up.

## Installation

Run `go get -u github.com/patrickmcnamara/down`.

Or you can download it from the GitHub release page if you like being weird.

## Usage

Run `down` after installation.
