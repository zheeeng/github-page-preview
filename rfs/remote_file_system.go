package rfs

import (
	"io/ioutil"
	"net/http"

	"github.com/github-page-preview/utils"
)

const proxyTarget = "https://raw.githubusercontent.com"

// RemoteFileSystem implements http.FileServer interface
type RemoteFileSystem interface {
	http.FileSystem
	SetReferer(string) RemoteFileSystem
	SetEndpointTransformer(endpointTransformer func(string) string) RemoteFileSystem
}

type remoteFileSystem struct {
	referer             string
	staticFolder        string
	endpointTransformer func(string) string
}

// NewRemoteFileSystem initializes RemoteFileSystem
func NewRemoteFileSystem(staticFolder string) RemoteFileSystem {
	return &remoteFileSystem{
		staticFolder: staticFolder,
	}
}

// SetEndpointTransformer set the endpoint transformer,
// which will be called before consuming the request endpoint.
func (rfs *remoteFileSystem) SetEndpointTransformer(endpointTransformer func(string) string) RemoteFileSystem {
	rfs.endpointTransformer = endpointTransformer
	return rfs
}

// ConfigureStatic set the static reading folder, when url matches seeking local files rule
func (rfs *remoteFileSystem) SetReferer(referer string) RemoteFileSystem {
	rfs.referer = referer
	return rfs
}

func (rfs *remoteFileSystem) Open(endpoint string) (RemoteFile, error) {
	if rfs.endpointTransformer != nil {
		endpoint = rfs.endpointTransformer(endpoint)
	}

	pc, err := utils.NewPathComponents(endpoint, rfs.referer)

	if err == utils.ErrNotMatchURLPattern {
		dir := http.Dir(rfs.staticFolder)
		return dir.Open(endpoint)
	} else if err != nil {
		return nil, err
	}

	resp, err := http.Get(proxyTarget + pc.Endpoint())

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	return NewRemoteFile(pc.GetFile(), data), nil
}
