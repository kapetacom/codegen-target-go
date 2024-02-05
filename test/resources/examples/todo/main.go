package main

import (
	"github.com/kapeta/todo/pkg/generated"
	"github.com/kapeta/todo/pkg/generated/auth"
	sdkgoconfig "github.com/kapetacom/sdk-go-config"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	generated.RegisterRouters(e, nil)
	config, err := sdkgoconfig.Init(".")
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
	e.Use(auth.AddJWTMiddleware(cfg)...)

	// Start the server and log if it fails
	e.Logger.Debug("Starting server on port " + port)
	e.Logger.Fatal(e.Start(host + ":" + port))
}
