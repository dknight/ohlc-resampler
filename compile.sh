#!/bin/bash

# Compiles ohlc-resampler for different arhitectures.

BINDIR="./bin"

mkdir -pv "$BINDIR"

GOOS=linux GOARCH=amd64 go build -o "$BINDIR/ohlc-resampler-amd64-linux"
GOOS=darwin GOARCH=amd64 go build -o "$BINDIR/ohlc-resampler-amd64-darwin"
GOOS=darwin GOARCH=arm64 go build -o "$BINDIR/ohlc-resampler-arm64-darwin"
GOOS=windows GOARCH=amd64 go build -o "$BINDIR/ohlc-resampler.exe"
