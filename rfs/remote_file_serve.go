package rfs

import (
	"net/http"

	"github.com/github-page-preview/utils"
)

// RemoteFileServe provides Start func
type RemoteFileServe interface {
	Start(res http.ResponseWriter, req *http.Request)
}

type remoteFileServe struct {
	remoteFileSystem RemoteFileSystem
}

// NewRemoteFileServe initializes RemoteFileServe
func NewRemoteFileServe(staticFolder string) RemoteFileServe {
	return &remoteFileServe{
		remoteFileSystem: NewRemoteFileSystem(staticFolder),
	}
}

func (rfsv *remoteFileServe) Start(res http.ResponseWriter, req *http.Request) {
	referer := req.Header.Get("Referer")
	if referer == "" {
		referer = req.Header.Get("referer")
	}

	rfsv.remoteFileSystem.SetReferer(referer)

	// Prevent the default redirection behavior caused by http.FileServer
	utils.PreventRedirection(req)

	http.FileServer(rfsv.remoteFileSystem).ServeHTTP(res, req)
}
