//#FILENAME:scripts/start-dev.sh:write-always:755
#!/bin/sh
if [ "$KAPETA_ENVIRONMENT_TYPE" = "docker" ]; then
  # In docker we want nodemon to exit on crash so that the container can be restarted
  go mod tidy
  go build -o app
  ./app
else
  go mod tidy
  go build -o app
  ./app
fi
