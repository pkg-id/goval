#!/usr/bin/env bash
set -e -o pipefail

echo "Execute go test"
if ! command -v tparse &> /dev/null ; then
    echo "tparse not installed or available in the PATH" >&2
    echo "please check https://pkg.go.dev/github.com/golang/tparse" >&2
    # Run go test on all packages.
   go test -race -count=1 -coverprofile=coverage.out -timeout 30s -v ./...
else
   # Run go test on all packages.
   go test -race -count=1 -coverprofile=coverage.out -timeout 30s -json ./... | tparse -all
fi
