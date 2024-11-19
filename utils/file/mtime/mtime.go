package mtime

import (
	"os"
	"time"
)

// layout defines a custom type representing a list of time format strings.
type layout []string

// ModTime represents a modification time given as a raw string.
type ModTime string

// NewModTime creates a new ModTime instance from the provided raw string.
// Parameters:
// - mtime: A string representing the modification time in a known format.
// Returns:
// - ModTime: A new instance of the ModTime struct.
func NewModTime(mtime string) ModTime {
	return ModTime(mtime)
}

// OlderOf checks if the ModTime instance represents a time older than the file's modification time.
// Parameters:
// - path: The file path whose modification time is to be compared.
// Returns:
// - bool: True if the ModTime instance is older than the file's modification time; otherwise, false.
// - error: Any error encountered during processing.
func (m ModTime) OlderOf(path string) (bool, error) {
	mtime, err := m.mtime()
	if err != nil {
		return false, err
	}
	ftime, err := m.ftime(path)
	if err != nil {
		return false, err
	}
	return mtime.After(ftime), err
}

// NewerOf checks if the ModTime instance represents a time newer than the file's modification time.
// Parameters:
// - path: The file path whose modification time is to be compared.
// Returns:
// - bool: True if the ModTime instance is newer than the file's modification time; otherwise, false.
// - error: Any error encountered during processing.
func (m ModTime) NewerOf(path string) (bool, error) {
	res, err := m.OlderOf(path)
	return !res, err
}

// mtime parses the raw modification time string into a time.Time object using predefined layouts.
// Returns:
// - time.Time: The parsed time.
// - error: An error if none of the layouts successfully parse the string.
func (m ModTime) mtime() (time.Time, error) {
	var (
		// Define a list of possible time formats.
		layouts = layout{
			time.DateTime,
			time.DateOnly,
			time.TimeOnly,
			time.Kitchen,
			time.Layout,
		}
	)

	// Attempt to parse the raw time string using each layout.
	for _, layout := range layouts {
		if mtime, err := time.Parse(layout, string(m)); err == nil {
			return mtime, nil
		}
	}

	// Return an error if none of the layouts match.
	return time.Time{}, NewInvalidTimeFormatError(string(m))
}

// ftime retrieves the modification time of the specified file.
// Parameters:
// - path: The file path to check.
// Returns:
// - time.Time: The file's modification time.
// - error: Any error encountered while accessing the file's information.
func (m ModTime) ftime(path string) (time.Time, error) {
	info, err := os.Stat(path)
	if err != nil {
		return time.Time{}, err
	}
	return info.ModTime(), nil
}
