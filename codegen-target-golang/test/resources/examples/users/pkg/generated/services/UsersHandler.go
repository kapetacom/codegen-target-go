//
// GENERATED SOURCE - DO NOT EDIT
//
package services

import (
	"github.com/labstack/echo/v4"
)

func CreateUser(ctx echo.Context) error {
	return ctx.JSON(200, map[string]interface{}{
		"message": "Hello from CreateUser",
	})
}

func GetUser(ctx echo.Context) error {
	return ctx.JSON(200, map[string]interface{}{
		"message": "Hello from GetUser",
	})
}

func DeleteUser(ctx echo.Context) error {
	return ctx.JSON(200, map[string]interface{}{
		"message": "Hello from DeleteUser",
	})
}
