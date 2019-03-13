package rfs

import (
	"net/http"
)

// LocalFileSystem implements http.FileServer interface
type LocalFileSystem interface {
	http.FileSystem
	SetEndpointTransformer(endpointTransformer func(string) string) LocalFileSystem
}

type localFileSystem struct {
	fileSystem          http.FileSystem
	endpointTransformer func(string) string
}

// NewLocalFileSystem initializes LocalFileSystem
func NewLocalFileSystem(staticFolder string) LocalFileSystem {
	return &localFileSystem{
		fileSystem: http.Dir(staticFolder),
	}
}

func (rfs *localFileSystem) Open(endpoint string) (RemoteFile, error) {
	if rfs.endpointTransformer != nil {
		endpoint = rfs.endpointTransformer(endpoint)
	}

	return rfs.fileSystem.Open(endpoint)
}

// SetEndpointTransformer set the endpoint transformer,
// which will be called before consuming the request endpoint.
func (rfs *localFileSystem) SetEndpointTransformer(endpointTransformer func(string) string) LocalFileSystem {
	rfs.endpointTransformer = endpointTransformer
	return rfs
}
