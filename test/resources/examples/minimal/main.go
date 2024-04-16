package main

import (
	kapeta "github.com/kapetacom/sdk-go-config"
	"github.com/kapetacom/sdk-go-rest-server/server"
)

func main() {
	serverInstance := server.NewWithDefaults()

	config, err := kapeta.Init(".")
	if err != nil {
		panic(err)
	}
	port, err := config.GetServerPort("rest")
	if err != nil {
		panic(err)
	}

	host, err := config.GetServerHost()
	if err != nil {
		panic(err)
	}

	// Start the server and log if it fails
	serverInstance.Logger.Debug("Starting server on port " + port)
	serverInstance.Logger.Fatal(serverInstance.Start(host + ":" + port))
}
