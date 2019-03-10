package rfs

import (
	"net/http"
	"os"
)

// RemoteFile alias to http.File interface
type RemoteFile = http.File

type remoteFile struct {
	at       int64
	name     string
	data     []byte
	length   int64
	fileInfo RemoteFileInfo
}

// NewRemoteFile initializes RemoteFile
func NewRemoteFile(name string, data []byte) RemoteFile {
	dataLength := int64(len(data))
	return &remoteFile{
		at:       0,
		name:     name,
		data:     data,
		length:   dataLength,
		fileInfo: NewRemoteFileInfo(name, dataLength),
	}
}

// Close pass
func (rf *remoteFile) Close() error { return nil }

// Stat pass
func (rf *remoteFile) Stat() (os.FileInfo, error) { return rf.fileInfo, nil }

// Readdir pass
func (rf *remoteFile) Readdir(count int) ([]os.FileInfo, error) {
	return []os.FileInfo{}, nil
}

// Read pass
func (rf *remoteFile) Read(b []byte) (int, error) {
	i := 0
	for rf.at < rf.length && i < len(b) {
		b[i] = rf.data[rf.at]
		i++
		rf.at++
	}
	return i, nil
}

// Seek pass
func (rf *remoteFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case 0:
		rf.at = offset
	case 1:
		rf.at += offset
	case 2:
		rf.at = rf.length + offset
	}
	return rf.at, nil
}
