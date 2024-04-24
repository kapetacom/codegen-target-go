package queues

import (
	"github.com/kapetacom/sdk-go-config/providers"
	"github.com/kapetacom/sdk-go-rabbitmq/rabbitmq"
)

type AnyEventsPublisher = rabbitmq.Publisher[any, map[string]any, string]

type AnyEventsPayload = rabbitmq.PublisherPayload[any, map[string]any, string]

func CreateAnyEventsPublisher(config providers.ConfigProvider) (*AnyEventsPublisher, error) {
	return rabbitmq.CreatePublisher[any, map[string]any, string](config, "anyEvents")
}
