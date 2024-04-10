// GENERATED SOURCE - DO NOT EDIT
package pubsub

import (
	"github.com/kapeta/todo/generated/entities"
	"github.com/kapetacom/sdk-go-config/providers"
	"github.com/kapetacom/sdk-go-google-pubsub/pubsub"
)

type UserPublisher = pubsub.Publisher[entities.User, map[string]string]

type UserPayload = pubsub.PublisherPayload[entities.User, map[string]string]

func CreateUserPublisher(config providers.ConfigProvider) (*UserPublisher, error) {
	return pubsub.CreatePublisher[entities.User, map[string]string](config, "pubsubpublisher")
}
