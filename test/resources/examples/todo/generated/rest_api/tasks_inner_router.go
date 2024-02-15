// GENERATED SOURCE - DO NOT EDIT
package rest

import (
	"fmt"

	generated "github.com/kapeta/todo/generated/services"
	"github.com/kapeta/todo/pkg/services"
	providers "github.com/kapetacom/sdk-go-config/providers"
	"github.com/kapetacom/sdk-go-rest-server/request"
	"github.com/labstack/echo/v4"
)

func CreateTasksInnerRouter(e *echo.Echo, cfg providers.ConfigProvider) error {
	routeHandler, err := services.NewTasksInnerHandler(cfg)
	if err != nil {
		return err
	}

	// Done like this to ensure interface compliance
	func(serviceInterface generated.TasksInnerInterface) {
		e.DELETE("/v2/tasks/:id", func(ctx echo.Context) error {
			var err error

			var id string
			if err = request.GetPathParams(ctx, "id", &id); err != nil {
				return ctx.String(400, fmt.Sprintf("bad request, unable to get path param id %v", err))
			}

			return serviceInterface.RemoveTask(ctx, id)
		})

		e.GET("/v2/tasks/:id", func(ctx echo.Context) error {
			var err error

			var id string
			if err = request.GetPathParams(ctx, "id", &id); err != nil {
				return ctx.String(400, fmt.Sprintf("bad request, unable to get path param id %v", err))
			}

			return serviceInterface.GetTask(ctx, id)
		})
	}(routeHandler)

	return nil
}
