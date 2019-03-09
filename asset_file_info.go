package main

import (
	"os"
	"time"
)

// AssetFileInfo pass
type AssetFileInfo struct {
	*AssetFile
}

// Name pass
func (s *AssetFileInfo) Name() string { return s.name }

// Size pass
func (s *AssetFileInfo) Size() int64 { return s.length }

// Mode pass
func (s *AssetFileInfo) Mode() os.FileMode { return os.ModeTemporary }

// ModTime pass
func (s *AssetFileInfo) ModTime() time.Time { return time.Time{} }

// IsDir pass
func (s *AssetFileInfo) IsDir() bool { return false }

// Sys pass
func (s *AssetFileInfo) Sys() interface{} { return nil }
