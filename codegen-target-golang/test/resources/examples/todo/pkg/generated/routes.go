package generated

import (
	"github.com/kapeta/todo/pkg/generated/rest"
	sdkgoconfig "github.com/kapetacom/golang-language-target/sdk-go-config/providers"
	"github.com/labstack/echo/v4"
)

func AddRoutes(e *echo.Echo, cfg sdkgoconfig.ConfigProvider) {
	rest.CreateTasksInnerRouter(e, cfg)

	rest.CreateTasksRouter(e, cfg)
}
