package utils

import (
	"fmt"
	"strings"
)

// ValidateRequired validates that a required field is not empty
func ValidateRequired(value, fieldName string) error {
	if strings.TrimSpace(value) == "" {
		return fmt.Errorf("%s is required", fieldName)
	}
	return nil
}

// TODO: Add specific validation functions as needed