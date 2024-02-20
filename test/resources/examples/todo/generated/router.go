package generated

import (
	rest "github.com/kapeta/todo/generated/rest_api"
	kapeta "github.com/kapetacom/sdk-go-config/providers"
	"github.com/kapetacom/sdk-go-rest-server/server"
)

func RegisterRouters(e *server.KapetaServer, cfg kapeta.ConfigProvider) error {
	var err error
	err = rest.CreateTasksInnerRouter(e, cfg)
	if err != nil {
		return err
	}

	err = rest.CreateTasksRouter(e, cfg)
	if err != nil {
		return err
	}
	return nil
}
