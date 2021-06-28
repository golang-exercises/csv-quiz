# csv-quiz

A quick and extensible command line quiz. Reference to [gophercise - Exercise 1](https://github.com/gophercises/quiz).

## Installation

Create a new directory in your `$GOPATH` and make sure that in your `$GOPATH` exists a `bin` directory.

```bash
# Project structure
$GOPATH
|
--- bin
--- src
     |
     --- <YOUR-PROJECT>
```

## Build

Install the `make` tool and run :

```bash
make build
```

## Run the program

```bash
$GOPATH/bin/quiz --file <path-to-csv> --timer <number> --shuffle <true/false>
```
