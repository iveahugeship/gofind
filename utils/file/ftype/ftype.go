package ftype

import (
	"io/fs"
)

// Predefined file types for easy identification of various file modes.
var (
	DevType       = NewFileType("DevType", fs.ModeDevice)
	CharDevType   = NewFileType("CharDevType", fs.ModeCharDevice)
	SymlinkType   = NewFileType("SymlinkType", fs.ModeSymlink)
	SocketType    = NewFileType("SocketType", fs.ModeSocket)
	NamedPipeType = NewFileType("NamedPipeType", fs.ModeNamedPipe)
	DirType       = NewFileType("DirType", fs.ModeDir)
	RegularType   = NewFileType("RegularType", 0)
)

// FileType represents a specific type of file, defined by a name and a file mode.
type FileType struct {
	name string
	mode fs.FileMode
}

// NewFileType creates a new instance of FileType with the provided name and mode.
// Parameters:
// - name: A string representing the name of the file type.
// - mode: The fs.FileMode corresponding to the file type.
// Returns:
// - FileType: A new FileType instance.
func NewFileType(name string, mode fs.FileMode) FileType {
	return FileType{
		name: name,
		mode: mode,
	}
}

// IsTypeOf checks if the provided file info matches the FileType's mode.
// Parameters:
// - info: An fs.FileInfo object containing information about a file.
// Returns:
// - bool: True if the file's mode matches the FileType's mode; otherwise, false.
func (f FileType) IsTypeOf(info fs.FileInfo) bool {
	return f.mode == info.Mode().Type()
}
