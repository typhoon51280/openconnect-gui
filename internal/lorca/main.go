package lorca

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/koltyakov/lorca"
	"github.com/typhoon51280/openconnect-gui/internal/ui"
)

func OpenWindow(wait bool, address string, args ...string) lorca.UI {

	url := ui.NewServer(address, true)

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
