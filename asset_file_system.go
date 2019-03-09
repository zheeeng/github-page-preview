package main

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
)

// AssetFileSystem implements http.FileServer
type AssetFileSystem struct{}

// Open pass
func (afs AssetFileSystem) Open(name string) (http.File, error) {
	if filepath.Ext(name) == "" {
		name += "/index.html"
	}

	resp, err := http.Get(proxyTarget + name)

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	af := AssetFile{
		at:     0,
		name:   name,
		data:   data,
		length: int64(len(data)),
	}

	return &af, nil
}
