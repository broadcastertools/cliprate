#!/usr/bin/env sh

set -e -v

go mod download

cat tools.go | grep _ | awk -F'"' '{print $2}' | xargs -t go install
