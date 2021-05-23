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

func NewServer(address string, start bool) string {
	if address == "" {
		address = "127.0.0.1:0"
	}
	http.Handle("/", CreateHandler())
	listener := CreateListener(address)
	if start {
		go Start(listener)
	}
	url := fmt.Sprintf("http://%s", listener.Addr())
	log.Printf("Listen on URL: %s", url)
	return url
}
