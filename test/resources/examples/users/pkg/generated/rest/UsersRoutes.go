//
// GENERATED SOURCE - DO NOT EDIT
//
package rest

import (
	"github.com/kapeta/todo/pkg/generated/services"
	sdkgoconfig "github.com/kapetacom/sdk-go-config/providers"
	"github.com/labstack/echo/v4"
)

func CreateUsersRouter(e *echo.Echo, cfg sdkgoconfig.ConfigProvider) {
	e.POST("/users/:id", services.CreateUser)

	e.GET("/users/:id", services.GetUser)

	e.DELETE("/users/:id", services.DeleteUser)
}
