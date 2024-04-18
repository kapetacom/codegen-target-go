// GENERATED SOURCE - DO NOT EDIT
package services

import (
	"github.com/kapeta/todo/generated/entities"
	"github.com/labstack/echo/v4"
)

// TasksInterface is an interface for TasksHandler
// It is used to define the methods that are implemented in the TasksHandler
type TasksInterface interface {
	GetData(ctx echo.Context, ids *[]string) error

	AddTask(ctx echo.Context, userId string, id string, task *entities.Task) error

	MarkAsDone(ctx echo.Context, id string) error
}
