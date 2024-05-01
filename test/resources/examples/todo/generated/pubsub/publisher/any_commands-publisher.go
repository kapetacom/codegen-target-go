// GENERATED SOURCE - DO NOT EDIT
package publisher

import (
	"github.com/kapetacom/sdk-go-config/providers"
	"github.com/kapetacom/sdk-go-google-pubsub/pubsub"
)

type AnyCommandsPublisher = pubsub.Publisher[any, map[string]string]

type AnyCommandsPayload = pubsub.PublisherPayload[any, map[string]string]

func CreateAnyCommandsPublisher(config providers.ConfigProvider) (*AnyCommandsPublisher, error) {
	return pubsub.CreatePublisher[any, map[string]string](config, "anyCommands")
}
