package depth

import (
	"path/filepath"
	"strings"
)

var (
	separator = string(filepath.Separator)
)

// Depth represents a utility to calculate and compare the depth of file paths relative to a specified root directory.
type Depth struct {
	root string
}

// NewDepth initializes a new Depth instance.
// Parameters:
// - root: The root directory to be used as the reference point.
// Returns:
// - Depth: An instance of the Depth struct with the cleaned and normalized root path.
func NewDepth(root string) Depth {
	return Depth{
		root: normalizePath(root),
	}
}

// UpperOf checks if the given path's depth is greater than or equal to the specified depth.
// Parameters:
// - path: The file path whose depth is to be compared.
// - n: The depth threshold for comparison.
// Returns:
// - bool: True if the path's depth is greater than or equal to `n`; otherwise, false.
func (d Depth) UpperOf(path string, n int) bool {
	if res, ok := d.depth(path); ok {
		return res >= n
	}
	return false
}

// LowerOf checks if the given path's depth is less than or equal to the specified depth.
// Parameters:
// - path: The file path whose depth is to be compared.
// - n: The depth threshold for comparison.
// Returns:
// - bool: True if the path's depth is less than or equal to `n`; otherwise, false.
func (d Depth) LowerOf(path string, n int) bool {
	if res, ok := d.depth(path); ok {
		return res <= n
	}
	return false
}

// depth calculates the relative depth of the given path with respect to the root directory.
// Parameters:
// - path: The file path for which the depth is calculated.
// Returns:
// - int: The depth of the path relative to the root. A positive value indicates levels deeper,
// a negative value indicates levels higher, and 0 indicates the same level.
// - bool: True if the path is valid and comparable to the root; otherwise, false.
func (d Depth) depth(path string) (int, bool) {
	path = normalizePath(path)
	diff := strings.TrimPrefix(path, d.root)
	count := strings.Count(diff, separator)

	// When path is lower than root or equal.
	if strings.HasPrefix(path, d.root) {
		return count, true
	}

	// When path is upper that root.
	if strings.HasPrefix(d.root, path) {
		return -count, true
	}

	// When path and root is completely different.
	return 0, false
}

// normalizePath normalizes a file path by cleaning it and ensuring it has a trailing separator.
// Parameters:
// - path: The file path to be normalized.
// Returns:
// - string: The cleaned and normalized file path.
func normalizePath(path string) string {
	return filepath.Clean(path) + separator
}
