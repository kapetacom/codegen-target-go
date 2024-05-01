package subscriber

import (
	"errors"
	"github.com/kapeta/users/generated/entities"
	"github.com/kapetacom/sdk-go-config/providers"
)

type UsersConsumer struct{}

func NewUsersConsumer(cfg providers.ConfigProvider) (*UsersConsumer, error) {
	return &UsersConsumer{}, nil
}

func (c *UsersConsumer) OnMessage(message entities.User, attributes map[string]string) error {
	if true {
		return errors.New("not implemented yet")
	}
	return nil
}
