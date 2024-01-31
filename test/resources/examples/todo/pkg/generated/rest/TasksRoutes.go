//
// GENERATED SOURCE - DO NOT EDIT
//
package rest

import (
	"github.com/kapeta/todo/pkg/generated/services"
	sdkgoconfig "github.com/kapetacom/sdk-go-config/providers"
	"github.com/labstack/echo/v4"
)

func CreateTasksRouter(e *echo.Echo, cfg sdkgoconfig.ConfigProvider) {
	e.POST("/tasks/:userid/:id", services.AddTask)

	e.POST("/tasks/:id/done", services.MarkAsDone)
}
