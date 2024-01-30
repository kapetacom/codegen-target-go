package clients

//
// GENERATED SOURCE - DO NOT EDIT
//
import (
	"encoding/json"
	"github.com/kapeta/todo/pkg/generated/entities"
	"github.com/kapeta/todo/pkg/sdk"
)

type Users interface {
	GetUserById(id string, metadata ...any) (*entities.User, error)

	DeleteUser(id string, metadata map[string]State, tags []string) error
}

type UsersClient struct {
	client *sdk.RestClient
}

// NewUsersClient creates new Users client
func NewUsersClient() Users {
	return &UsersClient{client: sdk.NewRestClient("Users", true)}
}

func (c *UsersClient) GetUserById(id string, metadata ...any) (*entities.User, error) {
	var result *entities.User
	resp, err := c.client.GET(c.client.ResolveURL("/users/%v", id))
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err

}

func (c *UsersClient) DeleteUser(id string, metadata map[string]State, tags []string) error {
	resp, err := c.client.DELETE(c.client.ResolveURL("/users/%v", id), metadata)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil

}
