//
// GENERATED SOURCE - DO NOT EDIT
//
package rest

import (
	"github.com/kapeta/todo/pkg/generated/services"
	sdkgoconfig "github.com/kapetacom/golang-language-target/sdk-go-config/providers"
	"github.com/labstack/echo/v4"
)

func CreateTasksInnerRouter(e *echo.Echo, cfg sdkgoconfig.ConfigProvider) {
	e.DELETE("/v2/tasks/:id", services.RemoveTask)

	e.GET("/v2/tasks/:id", services.GetTask)
}
