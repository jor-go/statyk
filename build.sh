#!/bin/bash

# Needs to run from this directory
mkdir -p output
go build -o output/statyk cmd/statyk/main.go
