#!/usr/bin/env bash

go mod tidy
go mod verify

echo "Running in Trace level"
ENV=dev SERVICE=exchange VERSION=1.0.0 LOG_LEVEL=trace go run examples/loglevel/main.go

echo

echo "Running in Debug level"
ENV=dev SERVICE=exchange VERSION=1.0.0 LOG_LEVEL=debug go run examples/loglevel/main.go

echo

echo "Running in Info level"
ENV=dev SERVICE=exchange VERSION=1.0.0 LOG_LEVEL=info go run examples/loglevel/main.go

echo

echo "Running in Warning level"
ENV=dev SERVICE=exchange VERSION=1.0.0 LOG_LEVEL=warning go run examples/loglevel/main.go

echo

echo "Running in Error level"
ENV=dev SERVICE=exchange VERSION=1.0.0 LOG_LEVEL=error go run examples/loglevel/main.go

