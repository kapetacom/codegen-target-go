package generated

import (
	"github.com/kapeta/todo/generated/rest_api"
	kapeta "github.com/kapetacom/sdk-go-config/providers"
	"github.com/labstack/echo/v4"
)

func RegisterRouters(e *echo.Echo, cfg kapeta.ConfigProvider) {
	rest.CreateHealth(e)

}
