#!/bin/sh

echo "Running tests..."
go test ./... || exit 1

# echo "Running linter..."
# golangci-lint run || exit 1