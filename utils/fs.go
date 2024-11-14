package utils

import (
	"path/filepath"
	"strings"
)

var (
	separator = string(filepath.Separator)
)

func CountDepth(root string, path string) int {
	root, path = filepath.Clean(root), filepath.Clean(path)
	diff := strings.TrimPrefix(path, root)
	return strings.Count(diff, separator) + 1
}
