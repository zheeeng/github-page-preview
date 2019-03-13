package main

import (
	"net/http"
	"os"

	"github.com/github-page-preview/rfs"
	"github.com/github-page-preview/utils"
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
	rfsv := rfs.NewRemoteFileServe(staticFolder, remoteHost)

	handler := utils.Compose(
		utils.IfElse(
			shouldCallMethodHandlerFunc,
			utils.Err(methodHandlerFunc),
			utils.Next(utils.EmptyHandlerFunc),
		),
		utils.Complete(rfsv.Start),
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
