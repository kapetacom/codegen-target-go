package main

import (
	"context"
	"github.com/kapeta/users/generated"
	"github.com/kapeta/users/generated/pubsub"
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

	err = generated.RegisterRouters(serverInstance, config)
	if err != nil {
		panic(err)
	}

	usersConsumer, err := pubsub.CreateUsersConsumer(config)
	if err != nil {
		panic(err)
	}
	go usersConsumer.ReceiveMessages(context.Background())

	anyEventsConsumer, err := pubsub.CreateAnyEventsConsumer(config)
	if err != nil {
		panic(err)
	}
	go anyEventsConsumer.ReceiveMessages(context.Background())

	// Start the server and log if it fails
	serverInstance.Logger.Debug("Starting server on port " + port)
	serverInstance.Logger.Fatal(serverInstance.Start(host + ":" + port))
}
