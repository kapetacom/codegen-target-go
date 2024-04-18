// GENERATED SOURCE - DO NOT EDIT
package rest

import (
	"github.com/kapeta/users/generated/entities"
	generated "github.com/kapeta/users/generated/services"
	"github.com/kapeta/users/pkg/services"
	providers "github.com/kapetacom/sdk-go-config/providers"
	"github.com/kapetacom/sdk-go-rest-server/request"
	"github.com/kapetacom/sdk-go-rest-server/server"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateUsersRouter(e *server.KapetaServer, cfg providers.ConfigProvider) error {
	routeHandler, err := services.NewUsersHandler(cfg)
	if err != nil {
		return err
	}

	// Done like this to ensure interface compliance
	func(serviceInterface generated.UsersInterface) {
		e.POST("/users/:id", func(ctx echo.Context) error {
			type RequestParameters struct {
				Id       string            `in:"path=id;required"`
				User     *entities.User    `in:"query=user;required"`
				Metadata map[string]string `in:"body=json"`
				Tags     *[]string         `in:"query=tags"`
			}
			params := &RequestParameters{}

			if err := request.GetRequestParameters(ctx.Request(), params); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return serviceInterface.CreateUser(ctx, params.Id, params.User, params.Metadata, params.Tags)
		})

		e.GET("/users", func(ctx echo.Context) error {
			type RequestParameters struct {
			}
			params := &RequestParameters{}

			if err := request.GetRequestParameters(ctx.Request(), params); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return serviceInterface.GetUsers(ctx)
		})

		e.GET("/users/:id", func(ctx echo.Context) error {
			type RequestParameters struct {
				Id       string `in:"path=id;required"`
				Metadata any    `in:"header=metadata;required"`
			}
			params := &RequestParameters{}

			if err := request.GetRequestParameters(ctx.Request(), params); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return serviceInterface.GetUser(ctx, params.Id, params.Metadata)
		})

		e.DELETE("/users/:id", func(ctx echo.Context) error {
			type RequestParameters struct {
				Id string `in:"path=id;required"`
			}
			params := &RequestParameters{}

			if err := request.GetRequestParameters(ctx.Request(), params); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return serviceInterface.DeleteUser(ctx, params.Id)
		})

		e.POST("/plan", func(ctx echo.Context) error {
			type RequestParameters struct {
				Type string `in:"query=type;required"`
				Body string `in:"body=json"`
			}
			params := &RequestParameters{}

			if err := request.GetRequestParameters(ctx.Request(), params); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return serviceInterface.HandlePlan(ctx, params.Type, params.Body)
		})
	}(routeHandler)

	return nil

}
