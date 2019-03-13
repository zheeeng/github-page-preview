package rfs

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// RemoteFileSystem implements http.FileServer interface
type RemoteFileSystem interface {
	http.FileSystem
	SetEndpointTransformer(endpointTransformer func(string) string) RemoteFileSystem
}

type remoteFileSystem struct {
	remoteHost          string
	endpointTransformer func(string) string
}

// NewRemoteFileSystem initializes RemoteFileSystem
func NewRemoteFileSystem(remoteHost string) RemoteFileSystem {
	return &remoteFileSystem{remoteHost: remoteHost}
}

func getName(endpoint string) string {
	chunks := strings.Split(endpoint, "/")
	return chunks[len(chunks)-1]
}

func (rfs *remoteFileSystem) Open(endpoint string) (RemoteFile, error) {
	if rfs.endpointTransformer != nil {
		endpoint = rfs.endpointTransformer(endpoint)
	}
	resp, err := http.Get(rfs.remoteHost + endpoint)

	if err != nil {
		return nil, errors.New("not found")
	}

	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	return NewRemoteFile(getName(endpoint), data), nil
}

// SetEndpointTransformer set the endpoint transformer,
// which will be called before consuming the request endpoint.
func (rfs *remoteFileSystem) SetEndpointTransformer(endpointTransformer func(string) string) RemoteFileSystem {
	rfs.endpointTransformer = endpointTransformer
	return rfs
}
