#!/usr/bin/env bash
jest --testMatch '<rootDir>/test/**/*.test.ts'
cd test/resources/examples/users
go mod tidy -v
go build -o users
# keep it clean
rm users
cd -
cd test/resources/examples/todo
go mod tidy -v
go build -o todo
# keep it clean
rm todo
cd -