package file

import (
	"os"
	"time"
)

type ModTime struct {
	t time.Time
}

func NewModTime(s string) (ModTime, error) {
	t, err := parseModTime(s)
	if err != nil {
		return ModTime{}, err
	}
	return ModTime{t: t}, nil
}

func (m ModTime) OlderOf(path string) (bool, error) {
	fileTime, err := getFileModTime(path)
	if err != nil {
		return false, err
	}
	return m.t.After(fileTime), err
}

func (m ModTime) NewerOf(path string) (bool, error) {
	fileTime, err := getFileModTime(path)
	if err != nil {
		return false, err
	}
	return m.t.Before(fileTime), nil
}

func parseModTime(s string) (time.Time, error) {
	var (
		err     error
		layouts = []string{
			time.DateTime,
			time.DateOnly,
			time.TimeOnly,
		}
	)

	for _, layout := range layouts {
		if t, err := time.Parse(layout, s); err == nil {
			return t, nil
		}
	}

	return time.Time{}, err
}

func getFileModTime(path string) (time.Time, error) {
	info, err := os.Stat(path)
	if err != nil {
		return time.Time{}, err
	}
	return info.ModTime(), nil
}
