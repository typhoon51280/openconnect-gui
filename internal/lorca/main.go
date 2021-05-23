package lorca

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/typhoon51280/openconnect-gui/internal/ui"
	"github.com/zserge/lorca"
)

func Listen(listener string) {

}

func OpenWindow(wait bool, args []string) lorca.UI {
	listener := "127.0.0.1:8888"
	url := fmt.Sprintf("http://%s", listener)
	log.Println(url)

	go ui.NewServer(listener)

	mainUI, err := lorca.New(url, "", 1024, 768, args...)
	if err != nil {
		log.Fatal(err)
	}
	mainUI.Bind("login", Login)
	if wait {
		Wait(mainUI)
	}
	return mainUI
}

func Wait(mainUI lorca.UI) {
	defer mainUI.Close()
	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigc:
	case <-mainUI.Done():
	}
}
