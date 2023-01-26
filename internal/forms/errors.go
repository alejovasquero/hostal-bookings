package forms

import (
	"fmt"
)

type errors map[string][]string

func (err *errors) AddError(field, message string) {
	(*err)[field] = append((*err)[field], message)
}

func (err *errors) GetError(field string) string {
	message, ok := (*err)[field]

	if !ok {
		return ""
	}

	return fmt.Sprint(message)
}
