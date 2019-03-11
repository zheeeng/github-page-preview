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
	SetPathTransformer(pathTransformer func(string) string) RemoteFileSystem
}

type remoteFileSystem struct {
	referer         string
	staticFolder    string
	pathTransformer func(string) string
}

// NewRemoteFileSystem initializes RemoteFileSystem
func NewRemoteFileSystem(staticFolder string) RemoteFileSystem {
	return &remoteFileSystem{
		staticFolder: staticFolder,
	}
}

// SetPathTransformer set the path transformer,
// which will be called before consuming the request path.
func (rfs *remoteFileSystem) SetPathTransformer(pathTransformer func(string) string) RemoteFileSystem {
	rfs.pathTransformer = pathTransformer
	return rfs
}

// ConfigureStatic set the static reading folder, when url matches seeking local files rule
func (rfs *remoteFileSystem) SetReferer(referer string) RemoteFileSystem {
	rfs.referer = referer
	return rfs
}

func (rfs *remoteFileSystem) Open(path string) (RemoteFile, error) {
	if rfs.pathTransformer != nil {
		path = rfs.pathTransformer(path)
	}

	pc, err := utils.NewPathComponents(path, rfs.referer)

	if err == utils.ErrNotMatchURLPattern {
		dir := http.Dir(rfs.staticFolder)
		return dir.Open(path)
	} else if err != nil {
		return nil, err
	}

	resp, err := http.Get(proxyTarget + pc.RequestPath())

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
