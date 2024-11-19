package mtime

import (
	"os"
	"time"
)

type layout []string

type ModTime struct {
	raw string
}

func NewModTime(mtime string) ModTime {
	return ModTime{raw: mtime}
}

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

func (m ModTime) NewerOf(path string) (bool, error) {
	res, err := m.OlderOf(path)
	return !res, err
}

func (m ModTime) mtime() (time.Time, error) {
	var (
		layouts = layout{
			time.DateTime,
			time.DateOnly,
			time.TimeOnly,
			time.Kitchen,
			time.Layout,
		}
	)

	for _, layout := range layouts {
		if mtime, err := time.Parse(layout, m.raw); err == nil {
			return mtime, nil
		}
	}

	return time.Time{}, NewInvalidTimeFormatError(m.raw)
}

func (m ModTime) ftime(path string) (time.Time, error) {
	info, err := os.Stat(path)
	if err != nil {
		return time.Time{}, err
	}
	return info.ModTime(), nil
}
