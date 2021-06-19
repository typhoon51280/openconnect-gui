package lorca

import (
	"log"
	"os"
	"os/signal"

	"github.com/koltyakov/lorca"
	"github.com/typhoon51280/openconnect-gui/internal/ui"
)

var mainWindow lorca.UI

type Connection struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
}

func OpenWindow(address string, embedded bool, wait bool, args ...string) lorca.UI {

	url := ui.NewServer(address, embedded)

	mainWindow, err := lorca.New("", "", 1024, 768, args...)
	if err != nil {
		log.Fatal(err)
	}
	mainWindow.Bind("Login", Login)
	mainWindow.Bind("Quit", Quit)
	mainWindow.Bind("Init", Init)
	mainWindow.Bind("Connect", Connect)
	mainWindow.Bind("Disconnect", Disconnect)
	mainWindow.Load(url)
	if wait {
		Wait(mainWindow)
	}
	return mainWindow
}

func Quit() {
	log.Println("Exit")
	mainWindow.Close()
}

func Connect() bool {
	log.Println("Connect")
	return true
}

func Disconnect() {
	log.Println("Disconnect")
}

func Init() []Connection {
	log.Println("Init")
	return []Connection{
		{
			Name:     "NFEC",
			Url:      "https://vpnssl.virtual.posteitaliane.it/NFEC",
			Username: "ntt3dr7@posteitaliane.it",
			Password: "Pegasus3",
		},
		{
			Name:     "SDP",
			Url:      "https://vpnssl.virtual.posteitaliane.it/sdpsvil2",
			Username: "ntt3dr7@posteitaliane.it",
			Password: "Pegasus3",
			Active:   true,
		},
	}
}

func Wait(window lorca.UI) {
	defer window.Close()
	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-window.Done():
	}
}
