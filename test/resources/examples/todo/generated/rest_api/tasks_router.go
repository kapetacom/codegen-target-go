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

func CreateTasksRouter(e *echo.Echo, cfg providers.ConfigProvider) error {
	routeHandler, err := services.NewTasksHandler()
	if err != nil {
		return err
	}
	handlerFunc := func(s genservices.TasksInterface) {
		e.POST("/tasks/:userid/:id", func(ctx echo.Context) error {
			var err error

			var userId = ctx.Param("userId")
			var id = ctx.Param("id")
			task := &entities.Task{}
			if _, err = request.GetBody(ctx, task); err != nil {
				return ctx.String(400, fmt.Sprintf("bad request, unable to unmarshal task %v", err))
			}
			return services.AddTask(ctx, userId, id, task)
		})

		e.POST("/tasks/:id/done", func(ctx echo.Context) error {
			var err error

			var id = ctx.Param("id")

			return services.MarkAsDone(ctx, id)
		})
	}
	handlerFunc(routeHandler)

	return nil
}
