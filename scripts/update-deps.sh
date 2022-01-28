#!/usr/bin/env sh

set -e

if [ ! -f "WORKSPACE" ]; then
    echo "Must be run from the root directory.";
    exit 1;
fi

set -v

export GOSUMDB=off
export GOPROXY=direct

go mod download -x

bazel run //:gazelle -- update-repos -from_file=./go.mod -to_macro "go_vendor.bzl%go_repositories"
