package main

import (
	"context"
	"github.com/kapeta/users/generated"
	"github.com/kapeta/users/generated/pubsub"
	kapeta "github.com/kapetacom/sdk-go-config"
	"github.com/kapetacom/sdk-go-rest-server/server"
)

func main() {
	e := server.NewWithDefaults()

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

	usersConsumer, err := pubsub.CreateUsersConsumer(config)
	if err != nil {
		panic(err)
	}
	go usersConsumer.ReceiveMessages(context.Background())

	// Start the server and log if it fails
	e.Logger.Debug("Starting server on port " + port)
	e.Logger.Fatal(e.Start(host + ":" + port))
}
