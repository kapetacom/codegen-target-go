// GENERATED SOURCE - DO NOT EDIT
package entities

// User - from Users block
type User struct {
	Idable
	Name string `json:"name" xml:"name" yaml:"name"`

	Email string `json:"email" xml:"email" yaml:"email"`

	State State `json:"state" xml:"state" yaml:"state"`
}
