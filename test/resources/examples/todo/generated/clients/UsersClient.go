package clients

//
// GENERATED SOURCE - DO NOT EDIT
//
import (
	"encoding/json"
	"github.com/kapeta/todo/generated/entities"
	"github.com/kapetacom/sdk-go-rest-client"
)

type Users interface {
	GetUserById(id string, metadata ...any) (*entities.User, error)

	DeleteUser(id string, tags *[]string) error
}

type UsersClient struct {
	client *client.RestClient
}

// NewUsersClient creates new Users client
func NewUsersClient() Users {
	return &UsersClient{client: client.NewRestClient("Users", true)}
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

func (c *UsersClient) DeleteUser(id string, tags *[]string) error {

	resp, err := c.client.DELETE(c.client.ResolveURL("/users/%v", id), client.QueryParameterRequestModifier(tags))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil

}
