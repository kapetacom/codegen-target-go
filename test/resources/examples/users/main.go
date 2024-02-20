package main

import (
	"github.com/kapeta/users/generated"
	kapeta "github.com/kapetacom/sdk-go-config"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

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

	err = generated.RegisterRouters(e, config)
	if err != nil {
		panic(err)
	}

	// Start the server and log if it fails
	e.Logger.Debug("Starting server on port " + port)
	e.Logger.Fatal(e.Start(host + ":" + port))
}
