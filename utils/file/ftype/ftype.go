package ftype

import (
	"io/fs"
)

var (
	DevType       = NewFileType("DevType", fs.ModeDevice)
	CharDevType   = NewFileType("CharDevType", fs.ModeCharDevice)
	SymlinkType   = NewFileType("SymlinkType", fs.ModeSymlink)
	SocketType    = NewFileType("SocketType", fs.ModeSocket)
	NamedPipeType = NewFileType("NamedPipeType", fs.ModeNamedPipe)
	DirType       = NewFileType("DirType", fs.ModeDir)
	RegularType   = NewFileType("RegularType", 0)
)

type FileType struct {
	name string
	mode fs.FileMode
}

func NewFileType(name string, mode fs.FileMode) FileType {
	return FileType{
		name: name,
		mode: mode,
	}
}

func (f FileType) IsTypeOf(info fs.FileInfo) bool {
	return f.mode == info.Mode().Type()
}
