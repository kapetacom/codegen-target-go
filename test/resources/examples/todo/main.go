package main

import (
	"context"
	"github.com/kapeta/todo/generated"
	"github.com/kapeta/todo/generated/auth"
	"github.com/kapeta/todo/generated/data"
	kapeta "github.com/kapetacom/sdk-go-config"
	"github.com/labstack/echo/v4"
	"log"
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

	e.Use(auth.AddJWTMiddleware(config)...)

	err = generated.RegisterRouters(e, config)
	if err != nil {
		panic(err)
	}

	dbTodo, closeTodo, err := data.NewTodo(config)
	if err != nil {
		log.Fatal(err)
	}
	defer closeTodo()

	// Start the server and log if it fails
	e.Logger.Debug("Starting server on port " + port)
	e.Logger.Fatal(e.Start(host + ":" + port))
}
