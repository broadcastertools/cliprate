#!/usr/bin/env sh

if [ ! -f WORKSPACE ]; then
    echo "This MUST be ran from the repository root!"
    exit 1
fi

# @todo move this to a Bazel script.

# Server code
rm -f core/api/api.go
oapi-codegen -package api -generate types,chi-server api.yaml  > core/api/api.go

# Client code

rm -rf ./app/src/api
npx openapi-typescript-codegen --input ./api.yaml --output ./app/src/api
