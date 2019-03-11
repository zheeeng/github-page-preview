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
		// Note: Here we setted a endpoint transformer, it will be called before consuming endpoint,
		// therefore we must call a reverse-direction transformer before feeding endpoint to consumer.
		// Look into Start func below, we called `utils.PreventRedirection(req)` for doing it.
		remoteFileSystem: NewRemoteFileSystem(staticFolder).SetEndpointTransformer(utils.RestoreHijacked),
	}
}

func (rfsv *remoteFileServe) Start(res http.ResponseWriter, req *http.Request) {
	referer := req.Header.Get("Referer")
	if referer == "" {
		referer = req.Header.Get("referer")
	}

	// Provide the request context in time
	rfsv.remoteFileSystem.SetReferer(referer)

	// Prevent the default redirection behavior caused by http.FileServer
	utils.PreventRedirection(req)

	http.FileServer(rfsv.remoteFileSystem).ServeHTTP(res, req)
}
