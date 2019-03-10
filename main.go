package main

import (
	"net/http"
	"os"

	"github.com/github-page-preview/utils"

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
	rfsv := rfs.NewRemoteFileServe("./public")

	http.HandleFunc("/", utils.HTTPHandlersCompose(
		methodHandler,
		utils.CompleteHandler(rfsv.Start),
	))

	http.ListenAndServe(getPort(), nil)
}

func methodHandler(res http.ResponseWriter, req *http.Request) utils.Cont {
	return func(next func(), complete func(), err func()) {
		if req.Method != http.MethodGet {
			http.Error(res, http.ErrBodyNotAllowed.Error(), http.StatusMethodNotAllowed)

			err()
		}

		next()
	}
}
