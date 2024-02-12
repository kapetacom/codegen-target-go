package generated

import (
	"github.com/kapeta/todo/generated/rest_api"
	kapeta "github.com/kapetacom/sdk-go-config/providers"
	"github.com/labstack/echo/v4"
)

func RegisterRouters(e *echo.Echo, cfg kapeta.ConfigProvider) error {
	var err error
	err = rest.CreateTasksInnerRouter(e, cfg)
	if err != nil {
		return err
	}

	err = rest.CreateTasksRouter(e, cfg)
	if err != nil {
		return err
	}

	err = rest.CreateHealth(e)
	if err != nil {
		return err
	}

	return nil
}
