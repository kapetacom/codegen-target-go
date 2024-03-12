//
// GENERATED SOURCE - DO NOT EDIT
//
package entities

type Status string

const _error Status = "error"
const success Status = "success"

func (s *Status) ToString() (string, error) {
	return string(*s), nil
}
func (s *Status) FromString(x string) error {
	*s = Status(x)
	return nil
}
