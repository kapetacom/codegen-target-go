//
// GENERATED SOURCE - DO NOT EDIT
//
package rest

import (
	genservices "github.com/kapeta/todo/pkg/generated/services"
	"github.com/kapeta/todo/pkg/services"
	sdkgoconfig "github.com/kapetacom/sdk-go-config/providers"
	"github.com/labstack/echo/v4"
)

func CreateTasksInnerRouter(e *echo.Echo, cfg sdkgoconfig.ConfigProvider) {
	services := &services.TasksInnerHandler{}
	handlerFunc := func(s genservices.TasksInnerInterface) {
		e.DELETE("/v2/tasks/:id", services.RemoveTask)

		e.GET("/v2/tasks/:id", services.GetTask)
	}
	handlerFunc(services)
}
