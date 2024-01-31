//
// GENERATED SOURCE - DO NOT EDIT
//
package services

import (
	"github.com/kapeta/todo/pkg/generated/entities"
	"github.com/labstack/echo/v4"
)

func RemoveTask(ctx echo.Context) error {

	return ctx.JSON(200, nil)
}

func GetTask(ctx echo.Context) error {

	return ctx.JSON(200, entities.Task{})
}
