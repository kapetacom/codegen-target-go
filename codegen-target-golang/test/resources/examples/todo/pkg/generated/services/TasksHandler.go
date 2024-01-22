//
// GENERATED SOURCE - DO NOT EDIT
//
package services

import (
	"github.com/labstack/echo/v4"
)

func AddTask(ctx echo.Context) error {
	return ctx.JSON(200, map[string]interface{}{
		"message": "Hello from AddTask",
	})
}

func MarkAsDone(ctx echo.Context) error {
	return ctx.JSON(200, map[string]interface{}{
		"message": "Hello from MarkAsDone",
	})
}
