//
// GENERATED SOURCE - DO NOT EDIT
//
package services

import (
	"github.com/kapeta/todo/pkg/generated/entities"
	"github.com/labstack/echo/v4"
)

func CreateUser(ctx echo.Context) error {
	_ = ctx.Param("user")
	_ = ctx.Param("tags")
	return ctx.JSON(200, entities.User{})
}

func GetUser(ctx echo.Context) error {

	return ctx.JSON(200, entities.User{})
}

func DeleteUser(ctx echo.Context) error {

	return ctx.JSON(200, nil)
}
