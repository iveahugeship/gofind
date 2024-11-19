package mtime

import (
	"fmt"
)

// InvalidTimeFormatError represents an error indicating that a provided time string is in an invalid format.
type InvalidTimeFormatError struct {
	value string
}

// NewInvalidTimeFormatError creates a new instance of InvalidTimeFormatError.
// Parameters:
// - value: The raw time string that caused the error.
// Returns:
// - InvalidTimeFormatError: An instance of the error with the provided value.
func NewInvalidTimeFormatError(value string) InvalidTimeFormatError {
	return InvalidTimeFormatError{
		value: value,
	}
}

// Error implements the error interface for InvalidTimeFormatError.
// Returns:
// - string: A formatted error message describing the invalid time format.
func (e InvalidTimeFormatError) Error() string {
	return fmt.Sprintf("invalid time format for '%s'", e.value)
}
