//
// GENERATED SOURCE - DO NOT EDIT
//
package entities

// A state
type State string

const OK State = "OK"
const NOT_OK State = "NOT_OK"

func (s *State) ToString() (string, error) {
	return string(*s), nil
}
func (s *State) FromString(x string) error {
	*s = State(x)
	return nil
}
