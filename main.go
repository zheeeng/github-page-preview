package main

import (
	"net/http"
	"os"

	"github.com/github-page-preview/fileserve"
	fph "github.com/golang-pkgs/functional-http-handler-funcs"
)

var (
	port         = "8090"
	staticFolder = "./public"
	remoteHost   = "https://raw.githubusercontent.com"
)

func init() {
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	if envStaticFolder := os.Getenv("STATIC_FOLDER"); envStaticFolder != "" {
		staticFolder = envStaticFolder
	}
}

func main() {
	fsv := fileserve.NewFileServe(staticFolder, remoteHost)

	handler := fph.Compose(
		fph.IfElse(
			shouldCallMethodHandlerFunc,
			fph.Err(methodHandlerFunc),
			fph.Next(fph.EmptyHandlerFunc),
		),
		fph.Complete(fsv.Start),
	)

	http.HandleFunc("/", handler)

	http.ListenAndServe(":"+port, nil)
}

func shouldCallMethodHandlerFunc(res http.ResponseWriter, req *http.Request) bool {
	return req.Method != http.MethodGet
}

func methodHandlerFunc(res http.ResponseWriter, req *http.Request) {
	http.Error(res, http.ErrBodyNotAllowed.Error(), http.StatusMethodNotAllowed)
}
