package domain

import (
	"bytes"
)

type FileInfo struct {
	Path     string
	Id       string
	Type     string
	FileData bytes.Buffer
}

func NewFile(path string, id string, fileType string, data bytes.Buffer) *FileInfo {
	return &FileInfo{Path: path, Id: id, Type: fileType, FileData: data}
}
