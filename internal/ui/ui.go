package ui

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed web/public
var web embed.FS

func CreateHandler() http.Handler {
	serverRoot, err := fs.Sub(web, "web/public")
	if err != nil {
		log.Fatal(err)
	}
	return http.FileServer(http.FS(serverRoot))
}

func NewServer(listener string) {
	http.Handle("/", CreateHandler())
	http.ListenAndServe(listener, nil)
}
