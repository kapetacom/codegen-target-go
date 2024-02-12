package queues

import (
	"github.com/kapeta/todo/generated/entities"
	"github.com/kapetacom/sdk-go-config/providers"
	"github.com/kapetacom/sdk-go-rabbitmq/rabbitmq"
)

type EventsPublisher = rabbitmq.Publisher[entities.User, map[string]any, string]

type EventsPayload = rabbitmq.PublisherPayload[entities.User, map[string]any, string]

func CreateEventsPublisher(config providers.ConfigProvider) (*EventsPublisher, error) {
	return rabbitmq.CreatePublisher[entities.User, map[string]any, string](config, "events")
}
