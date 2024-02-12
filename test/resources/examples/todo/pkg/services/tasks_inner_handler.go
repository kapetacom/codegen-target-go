package services

import (
	"github.com/kapeta/todo/generated/entities"
	providers "github.com/kapetacom/sdk-go-config/providers"
	"github.com/labstack/echo/v4"
)

type TasksInnerHandler struct {
}

func NewTasksInnerHandler(cfg providers.ConfigProvider) (*TasksInnerHandler, error) {
	return &TasksInnerHandler{}, nil
}

func (handler *TasksInnerHandler) RemoveTask(ctx echo.Context, id string) error {
	return ctx.JSON(200, nil)
}

func (handler *TasksInnerHandler) GetTask(ctx echo.Context, id string) error {
	return ctx.JSON(200, entities.Task{})
}
