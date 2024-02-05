package services

import (
	"github.com/kapeta/todo/pkg/generated/entities"
	"github.com/labstack/echo/v4"
)

type TasksInnerHandler struct {
}

func (u *TasksInnerHandler) RemoveTask(ctx echo.Context) error {

	return ctx.JSON(200, nil)
}

func (u *TasksInnerHandler) GetTask(ctx echo.Context) error {

	return ctx.JSON(200, entities.Task{})
}
