Advent of Code 2021
===================

This years advent of code. Solved with Go just to give me a refresher on the language.

#### Install Go

    curl -L https://go.dev/dl/go1.17.4.linux-amd64.tar.gz -o go1.17.4.linux-amd64.tar.gz
    sudo rm -rf /usr/local/go
    sudo tar -C /usr/local -xzf go1.17.4.linux-amd64.tar.gz
    rm go1.17.4.linux-amd64.tar.gz
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bash_profile
    . ~/.bash_profile
    go version
