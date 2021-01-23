package gooey

import "fmt"

func genMissingFieldError(field string) error {
	return fmt.Errorf("missing required field `%s`", field)
}

func invalidFieldValueError(field string) error {
	return fmt.Errorf("invalid field input `%s`", field)
}
