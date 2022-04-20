#!/bin/bash
export CGO_ENABLED=0
GOOS=linux GOARCH=amd64 go build -o pwdgen-linux64
GOOS=linux GOARCH=arm64 go build -o pwdgen-linux-arm64
GOOS=darwin GOARCH=amd64 go build -o pwdgen-darwin64
GOOS=darwin GOARCH=arm64 go build -o pwdgen-darwin-arm64
GOOS=windows GOARCH=amd64 go build -o pwdgen-windows64.exe
GOOS=windows GOARCH=386 go build -o pwdgen-windows32.exe
shasum -a 256 pwdgen-* > checksums.txt
