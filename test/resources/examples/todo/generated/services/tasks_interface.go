// GENERATED SOURCE - DO NOT EDIT
//
package services

import (
	"github.com/labstack/echo/v4"
)

// TasksInterface is an interface for TasksHandler
// It is used to define the methods that are implemented in the TasksHandler
type TasksInterface interface {
	AddTask(ctx echo.Context) error

	MarkAsDone(ctx echo.Context) error
}
