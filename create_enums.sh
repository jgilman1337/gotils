#!/bin/sh

## This script requires that `github.com/abice/go-enum` is installed and on your path. You can do this via running `go install github.com/abice/go-enum@latest`

go generate ./...
