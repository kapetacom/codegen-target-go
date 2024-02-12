//
// GENERATED SOURCE - DO NOT EDIT
//
package rest

import (
	"fmt"
	"github.com/kapeta/todo/generated/entities"
	genservices "github.com/kapeta/todo/generated/services"
	"github.com/kapeta/todo/pkg/services"
	providers "github.com/kapetacom/sdk-go-config/providers"
	"github.com/kapetacom/sdk-go-rest-server/request"
	"github.com/labstack/echo/v4"
)

func CreateTasksInnerRouter(e *echo.Echo, cfg providers.ConfigProvider) error {
	routeHandler, err := services.NewTasksInnerHandler()
	if err != nil {
		return err
	}
	handlerFunc := func(s genservices.TasksInnerInterface) {
		e.DELETE("/v2/tasks/:id", func(ctx echo.Context) error {
			var err error

			var id = ctx.Param("id")

			return services.RemoveTask(ctx, id)
		})

		e.GET("/v2/tasks/:id", func(ctx echo.Context) error {
			var err error

			var id = ctx.Param("id")

			return services.GetTask(ctx, id)
		})
	}
	handlerFunc(routeHandler)

	return nil
}
