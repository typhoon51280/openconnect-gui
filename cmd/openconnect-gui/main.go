package main

import (
	"flag"
	"log"
	"os"

	"github.com/typhoon51280/openconnect-gui/internal/electron"
	"github.com/typhoon51280/openconnect-gui/internal/lorca"
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

	var engine, address string
	var embedded, wait bool
	flag.StringVar(&engine, "engine", "lorca", "Engine type (lorca, muon, electron)")
	flag.StringVar(&address, "address", "127.0.0.1:0", "UI address")
	flag.BoolVar(&embedded, "embedded", true, "Start internal server")
	flag.BoolVar(&wait, "wait", true, "Wait for termination")
	flag.Parse()

	switch engine {
	case "lorca":
		lorca.OpenWindow(address, embedded, wait, os.Args...)
		break
	case "electron":
		electron.OpenWindow(true, address, os.Args...)
	}

	log.Println("Shutting down ...")
}
