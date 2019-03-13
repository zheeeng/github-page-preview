package fileserve

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

func (lfs *localFileSystem) Open(endpoint string) (RemoteFile, error) {
	if lfs.endpointTransformer != nil {
		endpoint = lfs.endpointTransformer(endpoint)
	}

	return lfs.fileSystem.Open(endpoint)
}

// SetEndpointTransformer set the endpoint transformer,
// which will be called before consuming the request endpoint.
func (lfs *localFileSystem) SetEndpointTransformer(endpointTransformer func(string) string) LocalFileSystem {
	lfs.endpointTransformer = endpointTransformer
	return lfs
}
