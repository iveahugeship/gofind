package utils

import (
	"time"
)

var (
	layouts = []string{
		time.DateTime,
		time.DateOnly,
		time.TimeOnly,
	}
)

func ParseDate(s string) (time.Time, error) {
	var err error
	for _, layout := range layouts {
		if d, err := time.Parse(layout, s); err == nil {
			return d, nil
		}
	}
	// Return zero value if parsing fails
	return time.Time{}, err
}
