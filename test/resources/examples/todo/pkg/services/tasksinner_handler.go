package services

import (
	"github.com/kapeta/todo/generated/entities"
	"github.com/labstack/echo/v4"
)

type TasksInnerHandler struct {
}

func (handler *TasksInnerHandler) RemoveTask(ctx echo.Context) error {

	return ctx.JSON(200, nil)
}

func (handler *TasksInnerHandler) GetTask(ctx echo.Context) error {

	return ctx.JSON(200, entities.Task{})
}
