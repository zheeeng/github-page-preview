package fileserve

import (
	"net/http"
	"net/url"

	"github.com/github-page-preview/utils"
)

// FileServe provides Start func
type FileServe interface {
	Start(res http.ResponseWriter, req *http.Request)
}

type fileServe struct {
	localFileHandler  http.Handler
	remoteFileHandler http.Handler
}

// NewFileServe initializes FileServe
func NewFileServe(staticFolder string, remoteHost string) FileServe {
	// Note: Here we setted endpoint transformers, they will be called before consuming endpoint,
	// therefore we must call a reverse-direction transformer before feeding endpoint to consumer.
	// Look into Start func below, we called `utils.PreventRedirection(req)` to do it.
	lfs := NewLocalFileSystem(staticFolder).SetEndpointTransformer(utils.RestoreHijacked)
	rfs := NewRemoteFileSystem(remoteHost).SetEndpointTransformer(utils.RestoreHijacked)

	return &fileServe{
		localFileHandler:  http.FileServer(lfs),
		remoteFileHandler: http.FileServer(rfs),
	}
}

func getReferer(req *http.Request) string {
	referer := req.Header.Get("Referer")
	if referer == "" {
		referer = req.Header.Get("referer")
	}
	URL, _ := url.Parse(referer)

	return URL.Path
}

func getEndpoint(req *http.Request) string {
	return req.URL.Path
}

func setEndpoint(req *http.Request, endpoint string) {
	req.URL.Path = endpoint
	utils.PreventRedirection(req)
}

func redirect(w http.ResponseWriter, r *http.Request, newPath string) {
	if q := r.URL.RawQuery; q != "" {
		newPath += "?" + q
	}
	w.Header().Set("Location", newPath)
	w.WriteHeader(http.StatusMovedPermanently)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func (fsv *fileServe) Start(res http.ResponseWriter, req *http.Request) {
	ec := utils.NewEndpointComponents(getEndpoint(req), getReferer(req))
	ep := ec.Endpoint()

	switch ec.GetEndpointType() {
	case utils.EndpointLocalAsset:
		setEndpoint(req, ep)
		fsv.localFileHandler.ServeHTTP(res, req)
	case utils.EndpointRemoteAsset:
		setEndpoint(req, ep)
		fsv.remoteFileHandler.ServeHTTP(res, req)
	case utils.EndpointRedirect:
		redirect(res, req, ep)
	case utils.EndpointNotFound:
		notFound(res, req)
	}
}
