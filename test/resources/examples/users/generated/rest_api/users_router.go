//
// GENERATED SOURCE - DO NOT EDIT
//
package rest

import (
	"fmt"
	"github.com/kapeta/users/generated/entities"
	generated "github.com/kapeta/users/generated/services"
	"github.com/kapeta/users/pkg/services"
	providers "github.com/kapetacom/sdk-go-config/providers"
	"github.com/kapetacom/sdk-go-rest-server/request"
	"github.com/kapetacom/sdk-go-rest-server/server"
	"github.com/labstack/echo/v4"
)

func CreateUsersRouter(e *server.KapetaServer, cfg providers.ConfigProvider) error {
	routeHandler, err := services.NewUsersHandler(cfg)
	if err != nil {
		return err
	}

	// Done like this to ensure interface compliance
	func(serviceInterface generated.UsersInterface) {
		e.POST("/users/:id", func(ctx echo.Context) error {
			var err error

			var user *entities.User
			if err = request.GetQueryParam(ctx, "user", user); err != nil {
				return ctx.String(400, fmt.Sprintf("bad request, unable to get query param user %v", err))
			}
			var tags *[]string
			if err = request.GetQueryParam(ctx, "tags", tags); err != nil {
				return ctx.String(400, fmt.Sprintf("bad request, unable to get query param tags %v", err))
			}
			var id string
			if err = request.GetPathParams(ctx, "id", &id); err != nil {
				return ctx.String(400, fmt.Sprintf("bad request, unable to get path param id %v", err))
			}
			var metadata map[string]string
			if err = request.GetBody(ctx, &metadata); err != nil {
				return ctx.String(400, fmt.Sprintf("bad request, unable to unmarshal metadata %v", err))
			}
			return serviceInterface.CreateUser(ctx, id, user, metadata, tags)
		})

		e.GET("/users", func(ctx echo.Context) error {

			return serviceInterface.GetUsers(ctx)
		})

		e.GET("/users/:id", func(ctx echo.Context) error {
			var err error
			var metadata any
			if err = request.GetHeaderParams(ctx, "metadata", &metadata); err != nil {
				return ctx.String(400, fmt.Sprintf("bad request, unable to get path param metadata %v", err))
			}

			var id string
			if err = request.GetPathParams(ctx, "id", &id); err != nil {
				return ctx.String(400, fmt.Sprintf("bad request, unable to get path param id %v", err))
			}

			return serviceInterface.GetUser(ctx, id, metadata)
		})

		e.DELETE("/users/:id", func(ctx echo.Context) error {
			var err error

			var id string
			if err = request.GetPathParams(ctx, "id", &id); err != nil {
				return ctx.String(400, fmt.Sprintf("bad request, unable to get path param id %v", err))
			}

			return serviceInterface.DeleteUser(ctx, id)
		})
	}(routeHandler)

	return nil
}
