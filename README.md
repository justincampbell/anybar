# anybar.go
[![Build Status](https://travis-ci.org/justincampbell/anybar.svg?branch=master)](https://travis-ci.org/justincampbell/anybar)
[![GoDoc](https://godoc.org/github.com/justincampbell/anybar?status.svg)](https://godoc.org/github.com/justincampbell/anybar)

> CLI and Go package for [AnyBar](https://github.com/tonsky/AnyBar)

## CLI Installation

### Go

    go get github.com/justincampbell/anybar/cmd/anybar

### Package

    wget -O anybar-latest.tar.gz https://github.com/justincampbell/anybar/archive/latest.tar.gz anybar
    tar -zxvf anybar-latest.tar.gz
    cd anybar-latest/
    make install

## Usage

    anybar red
    anybar green
    anybar blue
    anybar exclamation
    anybar -port 1234 question
