// GENERATED SOURCE - DO NOT EDIT
package subscriber

import (
	servicepubsub "github.com/kapeta/users/pkg/services/pubsub/subscriber"
	"github.com/kapetacom/sdk-go-config/providers"
	"github.com/kapetacom/sdk-go-google-pubsub/pubsub"
)

type AnyEventsConsumer = pubsub.Consumer[any]

type AnyEventsMessageHandler = pubsub.MessageHandler[any, map[string]string]

func CreateAnyEventsConsumer(config providers.ConfigProvider) (*AnyEventsConsumer, error) {
	consumer, err := servicepubsub.NewAnyEventsConsumer(config)
	if err != nil {
		return nil, err
	}
	return pubsub.CreateConsumer(config, "anyEvents", consumer.OnMessage)
}
