// GENERATED SOURCE - DO NOT EDIT
package pubsub

import (
	"github.com/kapeta/users/generated/entities"
	servicepubsub "github.com/kapeta/users/pkg/services/pubsub"
	"github.com/kapetacom/sdk-go-config/providers"
	"github.com/kapetacom/sdk-go-google-pubsub/pubsub"
)

type UserConsumer = pubsub.Consumer[entities.User]

type UserMessageHandler = pubsub.MessageHandler[entities.User, map[string]string]

func CreateUserConsumer(config providers.ConfigProvider) (*UserConsumer, error) {
	consumer, err := servicepubsub.NewUserConsumer(config)
	if err != nil {
		return nil, err
	}
	return pubsub.CreateConsumer(config, "pubsubsubscription", consumer.OnMessage)
}
