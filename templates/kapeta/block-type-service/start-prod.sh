//#FILENAME:scripts/start-prod.sh:write-always:755
#!/bin/sh
go mod tidy
go build -o app
./app