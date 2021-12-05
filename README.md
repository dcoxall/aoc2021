Advent of Code 2021
===================

This years advent of code. Solved with Go just to give me a refresher on the language.

## Structure

Each day is a submodule. The submodules (days) should take an `io.Reader` of the input
and calculate the correct result or return an error.

The main module brings all the days together and provides a basic CLI in which to execute
the solutions against the real input.

## Usage

Executing main.go will output the answer for the most recent solution but supports
re-execution of previous days. It will also output the time taken to calculate the
result.

    go run main.go -h
    go run main.go -d <DAY>

#### Install Go

    curl -L https://go.dev/dl/go1.17.4.linux-amd64.tar.gz -o go1.17.4.linux-amd64.tar.gz
    sudo rm -rf /usr/local/go
    sudo tar -C /usr/local -xzf go1.17.4.linux-amd64.tar.gz
    rm go1.17.4.linux-amd64.tar.gz
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bash_profile
    . ~/.bash_profile
    go version
