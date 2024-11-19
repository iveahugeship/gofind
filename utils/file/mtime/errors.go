package mtime

import (
	"fmt"
)

type InvalidTimeFormatError struct {
	value string
}

func NewInvalidTimeFormatError(value string) InvalidTimeFormatError {
	return InvalidTimeFormatError{
		value: value,
	}
}

func (e InvalidTimeFormatError) Error() string {
	return fmt.Sprintf("invalid time format for '%s'", e.value)
}
