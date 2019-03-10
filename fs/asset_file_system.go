package fs

import (
	"io/ioutil"
	"net/http"

	"github.com/github-page-preview/utils"
)

const proxyTarget = "https://raw.githubusercontent.com"

// AssetFileSystem alias to http.FileServer interface
type AssetFileSystem = http.FileSystem

type assetFileSystem struct {
	referer string
}

// NewAssetFileSystem initializes AssetFileSystem
func NewAssetFileSystem(referer string) AssetFileSystem {
	return &assetFileSystem{
		referer: referer,
	}
}

func (afs assetFileSystem) Open(name string) (AssetFile, error) {
	pc := utils.NewPathComponents(name, afs.referer)

	resp, err := http.Get(proxyTarget + pc.RequestPath())

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	return NewAssetFile(pc.GetFile(), data), nil
}
