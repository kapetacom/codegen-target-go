package main

import (
	"context"
	"github.com/kapeta/todo/generated"
	"github.com/kapeta/todo/generated/auth"
	"github.com/kapeta/todo/generated/data"
	kapeta "github.com/kapetacom/sdk-go-config"
	"github.com/kapetacom/sdk-go-rest-server/server"
	"log"
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

	serverInstance.Use(auth.AddJWTMiddleware(config)...)

	err = generated.RegisterRouters(serverInstance, config)
	if err != nil {
		panic(err)
	}

	dbTodo, closeTodo, err := data.NewTodo(config)
	if err != nil {
		log.Fatal(err)
	}
	defer closeTodo()
	err = dbTodo.Ping(context.Background(), nil) // Check if the database is alive
	if err != nil {
		log.Fatal(err)
	}

	// Start the server and log if it fails
	serverInstance.Logger.Debug("Starting server on port " + port)
	serverInstance.Logger.Fatal(serverInstance.Start(host + ":" + port))
}
