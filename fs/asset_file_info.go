package fs

import (
	"os"
	"time"
)

// AssetFileInfo alias to os.FileInfo
type AssetFileInfo = os.FileInfo

// assetFileInfo pass
type assetFileInfo struct {
	name   string
	length int64
}

// NewAssetFileInfo initializes AssetFileInfo
func NewAssetFileInfo(name string, length int64) AssetFileInfo {
	return &assetFileInfo{
		name:   name,
		length: length,
	}
}

// Name pass
func (s *assetFileInfo) Name() string { return s.name }

// Size pass
func (s *assetFileInfo) Size() int64 { return s.length }

// Mode pass
func (s *assetFileInfo) Mode() os.FileMode { return os.ModeTemporary }

// ModTime pass
func (s *assetFileInfo) ModTime() time.Time { return time.Time{} }

// IsDir pass
func (s *assetFileInfo) IsDir() bool { return false }

// Sys pass
func (s *assetFileInfo) Sys() interface{} { return nil }
