// GENERATED SOURCE - DO NOT EDIT
package subscriber

import (
	"github.com/kapeta/users/generated/entities"
	servicepubsub "github.com/kapeta/users/pkg/services/pubsub/subscriber"
	"github.com/kapetacom/sdk-go-config/providers"
	"github.com/kapetacom/sdk-go-google-pubsub/pubsub"
)

type UsersConsumer = pubsub.Consumer[entities.User]

type UsersMessageHandler = pubsub.MessageHandler[entities.User, map[string]string]

func CreateUsersConsumer(config providers.ConfigProvider) (*UsersConsumer, error) {
	consumer, err := servicepubsub.NewUsersConsumer(config)
	if err != nil {
		return nil, err
	}
	return pubsub.CreateConsumer(config, "users", consumer.OnMessage)
}
