package electron

import (
	"fmt"
	"log"
	"os"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"github.com/typhoon51280/openconnect-gui/internal/ui"
)

func OpenWindow(wait bool, address string, args ...string) {

	os.Unsetenv("ELECTRON_RUN_AS_NODE")

	url := ui.NewServer(address, true)

	// Initialize astilectron
	a, err := astilectron.New(log.New(os.Stderr, "", 0), astilectron.Options{
		AppName: "Openconnect-GUI",
	})
	if err != nil {
		log.Fatal(fmt.Errorf("main: creating astilectron failed: %w", err))
	}
	defer a.Close()

	// Handle signals
	a.HandleSignals()

	// Start astilectron
	if err = a.Start(); err != nil {
		log.Fatal(fmt.Errorf("main: starting astilectron failed: %w", err))
	}

	// New window
	var w *astilectron.Window
	if w, err = a.NewWindow(url, &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(700),
		Width:  astikit.IntPtr(700),
	}); err != nil {
		log.Fatal(fmt.Errorf("main: new window failed: %w", err))
	}

	// Create windows
	if err = w.Create(); err != nil {
		log.Fatal(fmt.Errorf("main: creating window failed: %w", err))
	}

	// Blocking pattern
	if wait {
		a.Wait()
	}
}
