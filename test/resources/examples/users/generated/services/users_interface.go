// GENERATED SOURCE - DO NOT EDIT
//
package services

import (
	"github.com/labstack/echo/v4"
)

// UsersInterface is an interface for UsersHandler
// It is used to define the methods that are implemented in the UsersHandler
type UsersInterface interface {
	CreateUser(ctx echo.Context) error

	GetUser(ctx echo.Context) error

	DeleteUser(ctx echo.Context) error
}
