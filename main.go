package main

import (
	"net/http"
	"os"
)

const (
	proxyTarget = "https://raw.githubusercontent.com"
	defaultPort = "8090"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	return ":" + port
}

func main() {
	http.HandleFunc("/", proxyServe)

	http.ListenAndServe(getPort(), nil)
}

func proxyServe(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, http.ErrBodyNotAllowed.Error(), http.StatusMethodNotAllowed)
	}

	http.FileServer(AssetFileSystem{}).ServeHTTP(res, req)
}
