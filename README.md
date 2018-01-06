# Slacking playground
A command line script that (given a webhook), reads messages from a file, and posts them to slack, written in
Python3, Go, and SCALA


## Setup

First setup an incoming webhook in slack and grab the url, (see http://bit.ly/2EapumJ for detailed instructions)

## Usage

### Python

#### Testing
    ./python/test.py

#### Running
    ./python/send.py --help
    ./python/send.py --webhook https://hooks.slack.com/services/XXXXX/XXXXX

### Go
    cd go
    go run send.go --help
    go run send.go --webhook https://hooks.slack.com/services/XXXXX/XXXXX

### SCALA
    TODO

### Futrue Enhancements

- Add tests for Go Implementation
- Allow Go script to run from anywhere
- Add tests for SCALA Implementation
- Implement SCALA script
