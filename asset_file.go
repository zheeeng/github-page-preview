package main

import (
	"net/http"
	"os"
)

// AssetFile alias to http.File interface
type AssetFile = http.File

type assetFile struct {
	at       int64
	name     string
	data     []byte
	length   int64
	fileInfo AssetFileInfo
}

// NewAssetFile initializes AssetFile
func NewAssetFile(name string, data []byte) AssetFile {
	dataLength := int64(len(data))
	return &assetFile{
		at:       0,
		name:     name,
		data:     data,
		length:   dataLength,
		fileInfo: NewAssetFileInfo(name, dataLength),
	}
}

// Close pass
func (af *assetFile) Close() error { return nil }

// Stat pass
func (af *assetFile) Stat() (os.FileInfo, error) { return af.fileInfo, nil }

// Readdir pass
func (af *assetFile) Readdir(count int) ([]os.FileInfo, error) {
	return []os.FileInfo{}, nil
}

// Read pass
func (af *assetFile) Read(b []byte) (int, error) {
	i := 0
	for af.at < af.length && i < len(b) {
		b[i] = af.data[af.at]
		i++
		af.at++
	}
	return i, nil
}

// Seek pass
func (af *assetFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case 0:
		af.at = offset
	case 1:
		af.at += offset
	case 2:
		af.at = af.length + offset
	}
	return af.at, nil
}
