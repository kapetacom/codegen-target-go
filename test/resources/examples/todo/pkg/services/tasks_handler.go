package services

import (
	"github.com/kapeta/todo/generated/entities"
	"github.com/labstack/echo/v4"
)

type TasksHandler struct {
}

func (handler *TasksHandler) AddTask(ctx echo.Context) error {

	return ctx.JSON(200, nil)
}

func (handler *TasksHandler) MarkAsDone(ctx echo.Context) error {

	return ctx.JSON(200, nil)
}
