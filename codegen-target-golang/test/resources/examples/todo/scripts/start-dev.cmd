@echo off
npm install
if "%KAPETA_ENVIRONMENT_TYPE%"=="docker" (
  rem In docker, we want nodemon to exit on crash so that the container can be restarted
  npm run start:dev:docker
) else (
  npm run start:dev
)
