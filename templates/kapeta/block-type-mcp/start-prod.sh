#!/bin/sh
# #FILENAME:start-prod.sh:create-only

# Build and run the production server
go build -o main .
./main