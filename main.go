package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path"
)

const (
	proxyTarget = "https://raw.githubusercontent.com"
	defaultPort = "8090"
)

func main() {
	http.HandleFunc("/", ProxyServe)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.ListenAndServe(":"+port, nil)
}

// ProxyServe starts a proxy service
func ProxyServe(rw http.ResponseWriter, req *http.Request) {
	u, _ := url.Parse(proxyTarget)

	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.URL.Scheme = u.Scheme
	req.URL.Host = u.Host
	req.URL.Path = path.Join(u.Path, req.URL.Path)

	httputil.NewSingleHostReverseProxy(u).ServeHTTP(rw, req)
}
