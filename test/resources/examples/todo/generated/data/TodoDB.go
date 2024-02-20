// GENERATED SOURCE - DO NOT EDIT
package data

import (
	"context"
	"github.com/kapetacom/sdk-go-config/providers"
	mongo "github.com/kapetacom/sdk-go-nosql-mongodb"
)

func NewTodo(config providers.ConfigProvider) (*mongo.MongoDB, func(), error) {
	// Create a new MongoDB client
	client, err := mongo.NewMongoDB(config, "todo")
	if err != nil {
		return nil, nil, err
	}
	closeFunc := func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}
	return client, closeFunc, nil
}
