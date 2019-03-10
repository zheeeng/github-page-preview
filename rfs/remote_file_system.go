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
}

type remoteFileSystem struct {
	referer      string
	staticFolder string
}

// NewRemoteFileSystem initializes RemoteFileSystem
func NewRemoteFileSystem(staticFolder string) RemoteFileSystem {
	return &remoteFileSystem{
		staticFolder: staticFolder,
	}
}

// ConfigureStatic set the static reading folder, when url matches seeking local files rule
func (rfs *remoteFileSystem) SetReferer(referer string) RemoteFileSystem {
	rfs.referer = referer
	return rfs
}

func (rfs *remoteFileSystem) Open(name string) (RemoteFile, error) {
	pc, err := utils.NewPathComponents(name, rfs.referer)

	if err == utils.ErrNotRecognize {
		dir := http.Dir(rfs.staticFolder)
		return dir.Open(name)
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
