package queues

import (
	"github.com/kapeta/todo/generated/entities"
	"github.com/kapetacom/sdk-go-config/providers"
	"github.com/kapetacom/sdk-go-rabbitmq/rabbitmq"
)

type EventsHandler = rabbitmq.MessageHandler[entities.User]

func CreateEventsConsumer(config providers.ConfigProvider, handler EventsHandler) (*rabbitmq.Consumer, error) {
	return rabbitmq.CreateConsumer[entities.User](config, "events", handler)
}
