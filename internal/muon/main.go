package muon

import (
	"log"

	"github.com/ImVexed/muon"
	"github.com/typhoon51280/openconnect-gui/internal/ui"
)

func OpenWindow(wait bool, args []string) {
	cfg := &muon.Config{
		Title:      "Hello, World!",
		Height:     500,
		Width:      500,
		Titled:     true,
		Resizeable: true,
	}
	m := muon.New(cfg, ui.CreateHandler())
	if err := m.Start(); err != nil {
		log.Fatal(err)
	}
}
