package clients

//
// GENERATED SOURCE - DO NOT EDIT
//
import (
	"encoding/json"
	"github.com/kapeta/todo/generated/entities"
	"github.com/kapetacom/sdk-go-rest-client"
)

type UsersInner interface {
	GetUserById(id string, metadata ...any) (*entities.User, error)

	DeleteUser(id string, tags *[]string) error
}

type UsersInnerClient struct {
	client *client.RestClient
}

// NewUsersInnerClient creates new UsersInner client
func NewUsersInnerClient() UsersInner {
	return &UsersInnerClient{client: client.NewRestClient("UsersInner", true)}
}

func (c *UsersInnerClient) GetUserById(id string, metadata ...any) (*entities.User, error) {
	var result *entities.User

	resp, err := c.client.GET(c.client.ResolveURL("/v2/users/%v", id))
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err

}

func (c *UsersInnerClient) DeleteUser(id string, tags *[]string) error {

	resp, err := c.client.DELETE(c.client.ResolveURL("/v2/users/%v", id), client.QueryParameterRequestModifier(tags))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil

}
