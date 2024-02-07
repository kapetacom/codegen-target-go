package generated

import (
	"github.com/kapeta/todo/generated/rest_api"
	sdkgoconfig "github.com/kapetacom/sdk-go-config/providers"
	"github.com/labstack/echo/v4"
)

func RegisterRouters(e *echo.Echo, cfg sdkgoconfig.ConfigProvider) {
	rest.CreateTasksInnerRouter(e, cfg)

	rest.CreateTasksRouter(e, cfg)
	rest.CreateHealth(e, cfg)
}
