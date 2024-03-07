package generated

import (
	rest "github.com/kapeta/todo/generated/rest_api"
	kapeta "github.com/kapetacom/sdk-go-config/providers"
	"github.com/kapetacom/sdk-go-rest-server/server"
)

func RegisterRouters(e *server.KapetaServer, cfg kapeta.ConfigProvider) error {
	if err := rest.CreateTasksInnerRouter(e, cfg); err != nil {
		return err
	}

	if err := rest.CreateTasksRouter(e, cfg); err != nil {
		return err
	}
	return nil
}
