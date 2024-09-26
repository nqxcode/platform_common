package validate

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// ValidationErrors represents a validation error
type ValidationErrors struct {
	Messages []string `json:"error_messages"`
}

func (v *ValidationErrors) addError(message string) {
	v.Messages = append(v.Messages, message)
}

// NewValidationErrors creates a new validation error
func NewValidationErrors(messages ...string) *ValidationErrors {
	return &ValidationErrors{
		Messages: messages,
	}
}

// Error returns the error message
func (v *ValidationErrors) Error() string {
	data, err := json.Marshal(v.Messages)
	if err != nil {
		return err.Error()
	}

	return string(data)
}

// IsValidationError checks if the error is a validation error
func IsValidationError(err error) bool {
	var ve *ValidationErrors
	return errors.As(err, &ve)
}
