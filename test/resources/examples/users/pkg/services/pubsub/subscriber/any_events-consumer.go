package subscriber

import (
	"errors"
	"github.com/kapetacom/sdk-go-config/providers"
)

type AnyEventsConsumer struct{}

func NewAnyEventsConsumer(cfg providers.ConfigProvider) (*AnyEventsConsumer, error) {
	return &AnyEventsConsumer{}, nil
}

func (c *AnyEventsConsumer) OnMessage(message any, attributes map[string]string) error {
	if true {
		return errors.New("not implemented yet")
	}
	return nil
}
