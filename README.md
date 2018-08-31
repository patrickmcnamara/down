# Down

This is a tool I made so I no longer have to type out `ping google.com` to check
if my internet is down.

It is very simple and *very* non-exhaustive. It simply makes a GET request to
Google. If that fails, it will say that your internet is down, else it will say
that your internet is up. It even does it with pretty colours.

## Installation

`go get -u github.com/patrickmcnamara/down`

Or you can download from the GitHub release tab if you want to be weird.

## Usage

`down`
