package utils

import (
	"encoding/base64"
	"net/http"
	"regexp"
	"strings"
)

const (
	indexStr           = "/index.html"
	indexPattern       = `/index\.html$`
	suffixSlashStr     = "/"
	suffixSlashPattern = "/$"
	delimiterStr       = "//"
	delimiterPattern   = `([^/])(//)([^/]+)$`
)

// from regex to string
var (
	indexFrom       = regexp.MustCompile(indexPattern)
	indexTo         = "/" + base64.URLEncoding.EncodeToString([]byte(indexPattern))
	suffixSlashFrom = regexp.MustCompile(suffixSlashPattern)
	suffixSlashTo   = "/" + base64.URLEncoding.EncodeToString([]byte(suffixSlashPattern))
	delimiterFrom   = regexp.MustCompile(delimiterPattern)
	delimiterTo     = "/" + base64.URLEncoding.EncodeToString([]byte(delimiterPattern))
)

func indexReplace(i string) (o string) {
	return string(indexFrom.ReplaceAll([]byte(i), []byte(indexTo)))
}
func indexRestore(i string) (o string) {
	return strings.Replace(i, indexTo, indexStr, -1)
}
func suffixSlashReplace(i string) (o string) {
	return string(suffixSlashFrom.ReplaceAll([]byte(i), []byte(suffixSlashTo)))
}
func suffixSlashRestore(i string) (o string) {
	return strings.Replace(i, suffixSlashTo, suffixSlashStr, -1)
}
func delimiterReplace(i string) (o string) {
	return string(delimiterFrom.ReplaceAll([]byte(i), []byte("$1"+delimiterTo+"$3")))
}
func delimiterRestore(i string) (o string) {
	return strings.Replace(i, delimiterTo, delimiterStr, -1)
}

// PreventRedirection hijacks endpoint, preventing 301 redirection by http.FileServer
func PreventRedirection(req *http.Request) {
	req.URL.Path = delimiterReplace(suffixSlashReplace(indexReplace(req.URL.Path)))
}

// RestoreHijacked restores the hijacked endpoint string before consuming
func RestoreHijacked(hijacked string) string {
	return indexRestore(suffixSlashRestore(delimiterRestore(hijacked)))
}
