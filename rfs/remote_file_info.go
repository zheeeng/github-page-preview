package rfs

import (
	"os"
	"time"
)

// RemoteFileInfo alias to os.FileInfo
type RemoteFileInfo = os.FileInfo

// remoteFileInfo pass
type remoteFileInfo struct {
	name   string
	length int64
}

// NewRemoteFileInfo initializes RemoteFileInfo
func NewRemoteFileInfo(name string, length int64) RemoteFileInfo {
	return &remoteFileInfo{
		name:   name,
		length: length,
	}
}

// Name pass
func (rfi *remoteFileInfo) Name() string { return rfi.name }

// Size pass
func (rfi *remoteFileInfo) Size() int64 { return rfi.length }

// Mode pass
func (rfi *remoteFileInfo) Mode() os.FileMode { return os.ModeTemporary }

// ModTime pass
func (rfi *remoteFileInfo) ModTime() time.Time { return time.Time{} }

// IsDir pass
func (rfi *remoteFileInfo) IsDir() bool { return false }

// Sys pass
func (rfi *remoteFileInfo) Sys() interface{} { return nil }
