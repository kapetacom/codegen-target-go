package pubsub

import (
	"errors"
	"github.com/kapeta/users/generated/entities"
	"github.com/kapetacom/sdk-go-config/providers"
)

type UserConsumer struct{}

func NewUserConsumer(cfg providers.ConfigProvider) (*UserConsumer, error) {
	return &UserConsumer{}, nil
}

func (c *UserConsumer) OnMessage(message entities.User, attributes map[string]string) error {
	if true {
		return errors.New("not implemented yet")
	}
	return nil
}
