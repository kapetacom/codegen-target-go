package queues

import (
	"github.com/kapetacom/sdk-go-config/providers"
	"github.com/kapetacom/sdk-go-rabbitmq/rabbitmq"
)

type AnyThingHandler = rabbitmq.MessageHandler[any]

func CreateAnyThingConsumer(config providers.ConfigProvider, handler AnyThingHandler) (*rabbitmq.Consumer, error) {
	return rabbitmq.CreateConsumer[any](config, "anyThing", handler)
}
