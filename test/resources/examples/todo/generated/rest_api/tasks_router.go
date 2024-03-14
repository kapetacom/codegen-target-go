//
// GENERATED SOURCE - DO NOT EDIT
//
package rest

import (
	"github.com/kapeta/todo/generated/entities"
	generated "github.com/kapeta/todo/generated/services"
	"github.com/kapeta/todo/pkg/services"
	providers "github.com/kapetacom/sdk-go-config/providers"
	"github.com/kapetacom/sdk-go-rest-server/request"
	"github.com/kapetacom/sdk-go-rest-server/server"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateTasksRouter(e *server.KapetaServer, cfg providers.ConfigProvider) error {
	routeHandler, err := services.NewTasksHandler(cfg)
	if err != nil {
		return err
	}

	// Done like this to ensure interface compliance
	func(serviceInterface generated.TasksInterface) {
		e.GET("/data", func(ctx echo.Context) error {
			type RequestParameters struct {
				Ids *[]string `in:"query=ids;required"`
			}
			params := &RequestParameters{}

			if err := request.GetRequestParameters(ctx.Request(), params); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return serviceInterface.GetData(ctx, params.Ids)
		})

		e.POST("/tasks/:userid/:id", func(ctx echo.Context) error {
			type RequestParameters struct {
				UserId string         `in:"path=userId;required"`
				Id     string         `in:"path=id;required"`
				Task   *entities.Task `in:"body=json"`
			}
			params := &RequestParameters{}

			if err := request.GetRequestParameters(ctx.Request(), params); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return serviceInterface.AddTask(ctx, params.UserId, params.Id, params.Task)
		})

		e.POST("/tasks/:id/done", func(ctx echo.Context) error {
			type RequestParameters struct {
				Id string `in:"path=id;required"`
			}
			params := &RequestParameters{}

			if err := request.GetRequestParameters(ctx.Request(), params); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return serviceInterface.MarkAsDone(ctx, params.Id)
		})
	}(routeHandler)

	return nil

}
