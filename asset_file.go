package main

import "os"

// AssetFile pass
type AssetFile struct {
	at     int64
	name   string
	data   []byte
	length int64
}

// Close pass
func (af *AssetFile) Close() error { return nil }

// Stat pass
func (af *AssetFile) Stat() (os.FileInfo, error) { return &AssetFileInfo{af}, nil }

// Readdir pass
func (af *AssetFile) Readdir(count int) ([]os.FileInfo, error) {
	return []os.FileInfo{}, nil
}

// Read pass
func (af *AssetFile) Read(b []byte) (int, error) {
	i := 0
	for af.at < af.length && i < len(b) {
		b[i] = af.data[af.at]
		i++
		af.at++
	}
	return i, nil
}

// Seek pass
func (af *AssetFile) Seek(offset int64, whence int) (int64, error) {
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
