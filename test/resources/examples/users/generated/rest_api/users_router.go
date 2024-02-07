//
// GENERATED SOURCE - DO NOT EDIT
//
package rest

import (
	"fmt"
	"github.com/kapeta/todo/generated/entities"
	genservices "github.com/kapeta/todo/generated/services"
	"github.com/kapeta/todo/pkg/services"
	sdkgoconfig "github.com/kapetacom/sdk-go-config/providers"
	"github.com/kapetacom/sdk-go-rest-server/request"
	"github.com/labstack/echo/v4"
)

func CreateUsersRouter(e *echo.Echo, cfg sdkgoconfig.ConfigProvider) {
	services := &services.UsersHandler{}
	handlerFunc := func(s genservices.UsersInterface) {
		e.POST("/users/:id", func(ctx echo.Context) error {
			var err error
			var user User
			if user, err = request.GetQueryParam(ctx, "user", &user); err != nil {
				return ctx.String(400, fmt.Sprintf("bad request, unable to get query param user %v", err))
			}
			var tags Set
			if tags, err = request.GetQueryParam(ctx, "tags", &tags); err != nil {
				return ctx.String(400, fmt.Sprintf("bad request, unable to get query param tags %v", err))
			}
			var id = ctx.Param("id")
			metadata := &map[string]string{}
			if body, err = request.GetBody(ctx, body); err != nil {
				return ctx.String(400, fmt.Sprintf("bad request, unable to unmarshal metadata %v", err))
			}
			return services.CreateUser(ctx, id, user, metadata, tags)
		})

		e.GET("/users/:id", func(ctx echo.Context) error {
			var err error

			var id = ctx.Param("id")

			return services.GetUser(ctx, id, metadata)
		})

		e.DELETE("/users/:id", func(ctx echo.Context) error {
			var err error

			var id = ctx.Param("id")

			return services.DeleteUser(ctx, id)
		})
	}
	handlerFunc(services)
}
