package services

import (
	"github.com/kapeta/todo/generated/entities"
	providers "github.com/kapetacom/sdk-go-config/providers"
	"github.com/labstack/echo/v4"
)

type TasksHandler struct {
}

func NewTasksHandler(cfg providers.ConfigProvider) (*TasksHandler, error) {
	return &TasksHandler{}, nil
}

func (handler *TasksHandler) AddTask(ctx echo.Context, userId string, id string, task *entities.Task) error {
	return ctx.JSON(200, nil)
}

func (handler *TasksHandler) MarkAsDone(ctx echo.Context, id string) error {
	return ctx.JSON(200, nil)
}
