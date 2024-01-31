#!/bin/sh
if [ "$KAPETA_ENVIRONMENT_TYPE" = "docker" ]; then
  # In docker we want nodemon to exit on crash so that the container can be restarted
  ulimit -n 2048
  go install github.com/cosmtrek/air@latest
  go mod tidy
  air -c .air.toml
else
  ulimit -n 2048
  go install github.com/cosmtrek/air@latest
  go mod tidy
  air -c .air.toml
fi
