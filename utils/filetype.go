package utils

import (
	"io/fs"
)

var (
	DevType       = NewFileType('b', fs.ModeDevice)
	CharDevType   = NewFileType('c', fs.ModeCharDevice)
	SymlinkType   = NewFileType('l', fs.ModeSymlink)
	SocketType    = NewFileType('s', fs.ModeSocket)
	NamedPipeType = NewFileType('p', fs.ModeNamedPipe)
	DirType       = NewFileType('d', fs.ModeDir)
	RegularType   = NewFileType('f', 0)
)

type FileType struct {
	String byte
	Mode   fs.FileMode
}

func NewFileType(s byte, m fs.FileMode) FileType {
	return FileType{
		String: s,
		Mode:   m,
	}
}
