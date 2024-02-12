package services

import (
	"github.com/kapeta/todo/generated/entities"
	providers "github.com/kapetacom/sdk-go-config/providers"
	"github.com/labstack/echo/v4"
)

type UsersHandler struct {
}

func NewUsersHandler(cfg providers.ConfigProvider) (*UsersHandler, error) {
	return &UsersHandler{}, nil
}

func (handler *UsersHandler) CreateUser(ctx echo.Context, id string, user *entities.User, metadata map[string]string, tags []string) error {
	return ctx.JSON(200, entities.User{})
}

func (handler *UsersHandler) GetUser(ctx echo.Context, id string, metadata any) error {
	return ctx.JSON(200, entities.User{})
}

func (handler *UsersHandler) DeleteUser(ctx echo.Context, id string) error {
	return ctx.JSON(200, nil)
}
