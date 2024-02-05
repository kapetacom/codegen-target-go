package services

import (
	"github.com/kapeta/todo/pkg/generated/entities"
	"github.com/labstack/echo/v4"
)

type TasksHandler struct {
}

func (u *TasksHandler) AddTask(ctx echo.Context) error {

	return ctx.JSON(200, nil)
}

func (u *TasksHandler) MarkAsDone(ctx echo.Context) error {

	return ctx.JSON(200, nil)
}
