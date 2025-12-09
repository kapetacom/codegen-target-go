#!/bin/sh
# #FILENAME:start-dev.sh:create-only

# Install air for live reloading if not present
if ! command -v air >/dev/null 2>&1; then
    echo "Installing air for live reloading..."
    go install github.com/cosmtrek/air@latest
fi

# Start the development server with live reloading
air