package generated

import (
	rest "github.com/kapeta/users/generated/rest_api"
	kapeta "github.com/kapetacom/sdk-go-config/providers"
	"github.com/labstack/echo/v4"
)

func RegisterRouters(e *echo.Echo, cfg kapeta.ConfigProvider) error {
	var err error
	err = rest.CreateUsersRouter(e, cfg)
	if err != nil {
		return err
	}

	err = rest.CreateHealth(e)
	if err != nil {
		return err
	}

	return nil
}
