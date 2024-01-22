//
// GENERATED SOURCE - DO NOT EDIT
//
package services

import (
	"github.com/labstack/echo/v4"
)

func RemoveTask(ctx echo.Context) error {
	return ctx.JSON(200, map[string]interface{}{
		"message": "Hello from RemoveTask",
	})
}

func GetTask(ctx echo.Context) error {
	return ctx.JSON(200, map[string]interface{}{
		"message": "Hello from GetTask",
	})
}
