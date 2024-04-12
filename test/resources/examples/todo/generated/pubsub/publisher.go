// GENERATED SOURCE - DO NOT EDIT
package pubsub

import (
	"github.com/kapeta/todo/generated/entities"
	"github.com/kapetacom/sdk-go-config/providers"
	"github.com/kapetacom/sdk-go-google-pubsub/pubsub"
)

type UsersPublisher = pubsub.Publisher[entities.User, map[string]string]

type UsersPayload = pubsub.PublisherPayload[entities.User, map[string]string]

func CreateUsersPublisher(config providers.ConfigProvider) (*UsersPublisher, error) {
	return pubsub.CreatePublisher[entities.User, map[string]string](config, "users")
}
