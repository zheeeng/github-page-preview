package main

import (
	"net/http"
	"os"

	"github.com/github-page-preview/rfs"
)

const defaultPort = "8090"

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	return ":" + port
}

func main() {
	http.HandleFunc("/", serve)

	http.ListenAndServe(getPort(), nil)
}

func serve(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, http.ErrBodyNotAllowed.Error(), http.StatusMethodNotAllowed)

		return
	}

	proxyServe(res, req)
}

func proxyServe(res http.ResponseWriter, req *http.Request) {
	referer := req.Header.Get("Referer")
	if referer == "" {
		referer = req.Header.Get("referer")
	}

	http.FileServer(rfs.NewRemoteFileSystem(referer)).ServeHTTP(res, req)
}
