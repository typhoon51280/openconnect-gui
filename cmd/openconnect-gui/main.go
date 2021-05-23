package main

import (
	"flag"
	"log"

	"github.com/typhoon51280/openconnect-gui/internal/electron"
	"github.com/typhoon51280/openconnect-gui/internal/lorca"
	"github.com/typhoon51280/openconnect-gui/internal/muon"
)

// func connect(url string, protocol string, tokenName string, tokenValue string) {
// 	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("sudo openconnect --protocol=%s --cookie=\"%s=%s\" \"%s\"", protocol, tokenName, tokenValue, url))
// 	log.Println(cmd.Args)
// 	cmd.Stdin = os.Stdin
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	// cmdReader, _ := cmd.StdoutPipe()
// 	// scanner := bufio.NewScanner(cmdReader)
// 	// done := make(chan bool)
// 	// go func() {
// 	// 	for scanner.Scan() {
// 	// 		fmt.Println(scanner.Text())
// 	// 	}
// 	// 	done <- true
// 	// }()
// 	cmd.Start()
// 	// <-done
// 	cmd.Wait()
// }

func main() {

	var engine string
	flag.StringVar(&engine, "engine", "electron", "Engine type (lorca, muon, electron)")

	switch engine {
	case "lorca":
		lorca.OpenWindow(true, nil)
		break
	case "muon":
		muon.OpenWindow(true, nil)
		break
	case "electron":
		electron.OpenWindow(true, nil)
	}

	log.Println("Shutting down ...")
}
