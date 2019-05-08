package utils

import (
	"github.com/fatih/structs"
	"github.com/google/uuid"
)

// ValidateError handles validate errors
type ValidateError struct {
	PropertyPath string `json:"propertyPath"`
	Message      string `json:"message"`
}

// ValidateUUID validate if the ob's fields is UUID
func ValidateUUID(ob interface{}, fields []string) []ValidateError {
	errors := []ValidateError{}
	for _, name := range fields {
		f := structs.Fields(ob)
		for _, field := range f {
			if name == field.Name() {
				_, err := uuid.Parse(field.Value().(string))
				if err != nil {
					errors = append(errors, ValidateError{
						PropertyPath: field.Name(),
						Message:      err.Error(),
					})
				}
			}
		}
	}
	return errors
}
