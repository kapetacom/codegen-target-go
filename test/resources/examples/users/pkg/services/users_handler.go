package services

import (
	"github.com/kapeta/todo/pkg/generated/entities"
	"github.com/labstack/echo/v4"
)

type UsersHandler struct {
}

func (u *UsersHandler) CreateUser(ctx echo.Context) error {
	_ = ctx.Param("user")
	_ = ctx.Param("tags")
	return ctx.JSON(200, entities.User{})
}

func (u *UsersHandler) GetUser(ctx echo.Context) error {

	return ctx.JSON(200, entities.User{})
}

func (u *UsersHandler) DeleteUser(ctx echo.Context) error {

	return ctx.JSON(200, nil)
}
