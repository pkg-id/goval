#!/usr/bin/env bash
set -e -o pipefail

echo "Execute go vet"
go vet ./...
