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

func CreateTasksRouter(e *echo.Echo, cfg sdkgoconfig.ConfigProvider) {
	services := &services.TasksHandler{}
	handlerFunc := func(s genservices.TasksInterface) {
		e.POST("/tasks/:userid/:id", services.AddTask)

		e.POST("/tasks/:id/done", services.MarkAsDone)
	}
	handlerFunc(services)
}
