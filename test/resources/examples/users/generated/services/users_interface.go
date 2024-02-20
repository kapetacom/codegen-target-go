// GENERATED SOURCE - DO NOT EDIT
//
package services

import (
	"github.com/kapeta/users/generated/entities"
	"github.com/labstack/echo/v4"
)

// UsersInterface is an interface for UsersHandler
// It is used to define the methods that are implemented in the UsersHandler
type UsersInterface interface {
	CreateUser(ctx echo.Context, id string, user *entities.User, metadata map[string]string, tags *[]string) error

	GetUser(ctx echo.Context, id string, metadata any) error

	DeleteUser(ctx echo.Context, id string) error
}
