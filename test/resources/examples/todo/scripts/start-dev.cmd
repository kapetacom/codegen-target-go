@echo off
if "%KAPETA_ENVIRONMENT_TYPE%"=="docker" (
  rem In docker, we want nodemon to exit on crash so that the container can be restarted
  go mod tidy
  go build -o app
  ./app
) else (
  go mod tidy
  go build -o app
  ./app
)
