// GENERATED SOURCE - DO NOT EDIT
package rest

import (
	generated "github.com/kapeta/todo/generated/services"
	"github.com/kapeta/todo/pkg/services"
	providers "github.com/kapetacom/sdk-go-config/providers"
	"github.com/kapetacom/sdk-go-rest-server/request"
	"github.com/kapetacom/sdk-go-rest-server/server"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateTasksInnerRouter(e *server.KapetaServer, cfg providers.ConfigProvider) error {
	routeHandler, err := services.NewTasksInnerHandler(cfg)
	if err != nil {
		return err
	}

	// Done like this to ensure interface compliance
	func(serviceInterface generated.TasksInnerInterface) {
		e.DELETE("/v2/tasks/:id", func(ctx echo.Context) error {
			type RequestParameters struct {
				Id string `in:"path=id;required"`
			}
			params := RequestParameters{}

			if err := request.FillStruct(ctx, &params); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return serviceInterface.RemoveTask(ctx, params.Id)
		})

		e.GET("/v2/tasks/:id", func(ctx echo.Context) error {
			type RequestParameters struct {
				Id string `in:"path=id;required"`
			}
			params := RequestParameters{}

			if err := request.FillStruct(ctx, &params); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return serviceInterface.GetTask(ctx, params.Id)
		})
	}(routeHandler)

	return nil

}
