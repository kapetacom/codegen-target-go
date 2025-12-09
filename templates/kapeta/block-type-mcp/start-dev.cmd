@echo off
REM #FILENAME:start-dev.cmd:create-only

REM Check if air is installed
where air >nul 2>&1
if %errorlevel% neq 0 (
    echo Installing air for live reloading...
    go install github.com/cosmtrek/air@latest
)

REM Start the development server with live reloading
air