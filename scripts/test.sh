#!/usr/bin/env bash
jest --testMatch '<rootDir>/test/**/*.test.ts'
cd test/resources/examples/users
go mod tidy -v
go build ./...
cd -
cd test/resources/examples/todo
go mod tidy -v
go build ./...
cd -