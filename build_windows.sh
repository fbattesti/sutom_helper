#!/bin/bash 

export GOOS=windows
export GOARCH=amd64

go build -o sutom.exe 

unset GOOS
unset GOARCH
