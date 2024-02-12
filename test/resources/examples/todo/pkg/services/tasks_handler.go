package services

import (
	"github.com/kapeta/todo/generated/entities"
	"github.com/labstack/echo/v4"
)

type TasksHandler struct {
}

func NewTasksHandler() (*TasksHandler, error) {
	return &TasksHandler{}, nil
}

func (handler *TasksHandler) AddTask(ctx echo.Context, userId string, id string, task *entities.Task) error {
	return ctx.JSON(200, nil)
}

func (handler *TasksHandler) MarkAsDone(ctx echo.Context, id string) error {
	return ctx.JSON(200, nil)
}
