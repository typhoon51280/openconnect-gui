package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/zserge/lorca"
)

//go:embed frontend/public
var frontend_public embed.FS

func main() {
	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}

	// launch Chrome Instance (lorca)
	ui, err := lorca.New("", "", 480, 320, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	// listen on localhost random socket
	server, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	// serve frontend
	serverRoot, err := fs.Sub(frontend_public, "frontend/public")
	if err != nil {
		log.Fatal(err)
	}
	go http.Serve(server, http.FileServer(http.FS(serverRoot)))

	url := fmt.Sprintf("http://%s", server.Addr())
	log.Println(url)
	ui.Load(url)

	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigc:
	case <-ui.Done():
	}
	log.Println("Shutting down ...")
}
