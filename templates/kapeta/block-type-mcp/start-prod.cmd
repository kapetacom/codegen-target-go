@echo off
REM #FILENAME:start-prod.cmd:create-only

REM Build and run the production server
go build -o main.exe .
main.exe