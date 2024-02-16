// GENERATED SOURCE - DO NOT EDIT
//
package services

import (
	"github.com/labstack/echo/v4"
)

// TasksInnerInterface is an interface for TasksInnerHandler
// It is used to define the methods that are implemented in the TasksInnerHandler
type TasksInnerInterface interface {
	RemoveTask(ctx echo.Context, id string) error

	GetTask(ctx echo.Context, id string) error
}
