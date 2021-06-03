package ui

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net"
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

func CreateListener(address string) net.Listener {
	server, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	return server
}

func Start(listener net.Listener) {
	defer listener.Close()
	http.Serve(listener, nil)
}

func NewServer(address string, embedded bool) string {
	if embedded {
		http.Handle("/", CreateHandler())
		listener := CreateListener(address)
		address = listener.Addr().String()
		go Start(listener)
		log.Printf("Listen on URL: %s", address)
	}
	url := fmt.Sprintf("http://%s", address)
	return url
}
