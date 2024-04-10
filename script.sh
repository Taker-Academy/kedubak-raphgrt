#!/bin/bash

# Set environment variables for cross-compiling to ARM64v8
export GOOS=darwin
export GOARCH=arm64
export CGO_ENABLED=0

# Download dependencies
go mod download

# Cross-compile your application
go build -o KeDuBack .

# Run your tests
go test -v ./...

# Unset environment variables
unset GOOS
unset GOARCH
unset CGO_ENABLED