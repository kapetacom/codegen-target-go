#!/usr/bin/env bash
if ! jest --testMatch '<rootDir>/test/**/*.test.ts'; then
  exit $RESULT
fi
cd test/resources/examples/users
go mod tidy -v
go build ./...
cd -
cd test/resources/examples/todo
go mod tidy -v
go build ./...
cd -