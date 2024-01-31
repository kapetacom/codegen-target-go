#!/bin/sh
go mod tidy
go build -o app
./app