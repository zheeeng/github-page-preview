package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

const (
	rootName    = "gpr"
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
	u, _ := url.Parse(proxyTarget)

	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = u.Host
	req.URL.Scheme = u.Scheme
	req.URL.Host = u.Host

	httputil.NewSingleHostReverseProxy(u).ServeHTTP(res, req)
}
