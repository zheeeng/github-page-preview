package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/github-page-preview/fs"
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

	if req.URL.Path == "/" {
		welcomeserve(res, req)

		return
	}

	proxyServe(res, req)
}

func welcomeserve(res http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("public/index.html"))
	tmpl.Execute(res, nil)
}

func proxyServe(res http.ResponseWriter, req *http.Request) {
	referer := req.Header.Get("Referer")
	if referer == "" {
		referer = req.Header.Get("referer")
	}

	http.FileServer(fs.NewAssetFileSystem(referer)).ServeHTTP(res, req)
}
