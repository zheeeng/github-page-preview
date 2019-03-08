package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
)

const proxyTarget = "https://raw.githubusercontent.com"

func main() {
	http.HandleFunc("/", ProxyServe)

	http.ListenAndServe(":8090", nil)
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
