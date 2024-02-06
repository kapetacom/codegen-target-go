package services

import (
	"github.com/kapeta/todo/generated/entities"
	"github.com/labstack/echo/v4"
)

type UsersHandler struct {
}

func (handler *UsersHandler) CreateUser(ctx echo.Context) error {
	_ = ctx.Param("user")
	_ = ctx.Param("tags")
	return ctx.JSON(200, entities.User{})
}

func (handler *UsersHandler) GetUser(ctx echo.Context) error {

	return ctx.JSON(200, entities.User{})
}

func (handler *UsersHandler) DeleteUser(ctx echo.Context) error {

	return ctx.JSON(200, nil)
}
