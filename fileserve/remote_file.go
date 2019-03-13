package fileserve

import (
	"net/http"
)

// RemoteFile alias to http.File interface
type RemoteFile = http.File

type remoteFile struct {
	cursor   int64
	name     string
	data     []byte
	length   int64
	fileInfo RemoteFileInfo
}

// NewRemoteFile initializes RemoteFile
func NewRemoteFile(name string, data []byte) RemoteFile {
	dataLength := int64(len(data))
	return &remoteFile{
		cursor:   0,
		name:     name,
		data:     data,
		length:   dataLength,
		fileInfo: NewRemoteFileInfo(name, dataLength),
	}
}

// Close pass
func (rf *remoteFile) Close() error { return nil }

// Stat pass
func (rf *remoteFile) Stat() (RemoteFileInfo, error) { return rf.fileInfo, nil }

// Readdir pass
func (rf *remoteFile) Readdir(count int) ([]RemoteFileInfo, error) {
	return []RemoteFileInfo{}, nil
}

// Read pass
func (rf *remoteFile) Read(b []byte) (int, error) {
	i := 0
	for rf.cursor < rf.length && i < len(b) {
		b[i] = rf.data[rf.cursor]
		i++
		rf.cursor++
	}
	return i, nil
}

// Seek pass
func (rf *remoteFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case 0:
		rf.cursor = offset
	case 1:
		rf.cursor += offset
	case 2:
		rf.cursor = rf.length + offset
	}
	return rf.cursor, nil
}
