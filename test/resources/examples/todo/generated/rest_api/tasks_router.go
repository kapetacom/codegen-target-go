//
// GENERATED SOURCE - DO NOT EDIT
//
package rest

import (
	"fmt"
	"github.com/kapeta/todo/generated/entities"
	generated "github.com/kapeta/todo/generated/services"
	"github.com/kapeta/todo/pkg/services"
	providers "github.com/kapetacom/sdk-go-config/providers"
	"github.com/kapetacom/sdk-go-rest-server/request"
	"github.com/kapetacom/sdk-go-rest-server/server"
	"github.com/labstack/echo/v4"
)

func CreateTasksRouter(e *server.KapetaServer, cfg providers.ConfigProvider) error {
	routeHandler, err := services.NewTasksHandler(cfg)
	if err != nil {
		return err
	}

	// Done like this to ensure interface compliance
	func(serviceInterface generated.TasksInterface) {
		e.POST("/tasks/:userid/:id", func(ctx echo.Context) error {
			var err error

			var userId string
			if err = request.GetPathParams(ctx, "userId", &userId); err != nil {
				return ctx.String(400, fmt.Sprintf("bad request, unable to get path param userId %v", err))
			}
			var id string
			if err = request.GetPathParams(ctx, "id", &id); err != nil {
				return ctx.String(400, fmt.Sprintf("bad request, unable to get path param id %v", err))
			}
			var task *entities.Task
			if err = request.GetBody(ctx, task); err != nil {
				return ctx.String(400, fmt.Sprintf("bad request, unable to unmarshal task %v", err))
			}
			return serviceInterface.AddTask(ctx, userId, id, task)
		})

		e.POST("/tasks/:id/done", func(ctx echo.Context) error {
			var err error

			var id string
			if err = request.GetPathParams(ctx, "id", &id); err != nil {
				return ctx.String(400, fmt.Sprintf("bad request, unable to get path param id %v", err))
			}

			return serviceInterface.MarkAsDone(ctx, id)
		})
	}(routeHandler)

	return nil
}
