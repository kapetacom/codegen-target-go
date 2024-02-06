//
// GENERATED SOURCE - DO NOT EDIT
//
package rest

import (
	genservices "github.com/kapeta/todo/generated/services"
	"github.com/kapeta/todo/pkg/services"
	sdkgoconfig "github.com/kapetacom/sdk-go-config/providers"
	"github.com/labstack/echo/v4"
)

func CreateUsersRouter(e *echo.Echo, cfg sdkgoconfig.ConfigProvider) {
	services := &services.UsersHandler{}
	handlerFunc := func(s genservices.UsersInterface) {
		e.POST("/users/:id", services.CreateUser)

		e.GET("/users/:id", services.GetUser)

		e.DELETE("/users/:id", services.DeleteUser)
	}
	handlerFunc(services)
}
