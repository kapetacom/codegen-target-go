//#FILENAME:scripts/start-prod.cmd:write-always:644
@echo off
go mod tidy
go build -o app
./app