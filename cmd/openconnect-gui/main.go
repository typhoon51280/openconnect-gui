package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/zserge/lorca"
)

//go:embed frontend/public/index.html frontend/public/build/* frontend/public/assets/*
var frontend_public embed.FS

type CookieInterval struct {
	running chan bool
	cookie  string
}

func login(loginUrl string, email string, password string) string {
	log.Printf("login %s\n", loginUrl)
	args := []string{}
	ui, err := lorca.New(loginUrl, "", 1000, 600, args...)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("login opened")
	defer ui.Close()
	cookieInterval := getCookies(ui, loginUrl, email, password)
	// <-ui.Done()
	<-cookieInterval.running
	log.Println("login closed")
	// stop(cookieInterval)
	// log.Println("cookieInterval stopped")
	// if cookieInterval.cookie != "" {
	// 	go connect(loginUrl, "nc", "DSID", cookieInterval.cookie)
	// }
	return cookieInterval.cookie
}

func stop(cookieInterval *CookieInterval) {
	cookieInterval.running <- true
	close(cookieInterval.running)
}

func getCookies(ui lorca.UI, loginUrl string, email string, password string) *CookieInterval {
	ui.Eval("document.getCookies")
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan bool)
	cookieInterval := &CookieInterval{
		running: quit,
		cookie:  "",
	}
	go func() {
		urlObj, _ := url.Parse(loginUrl)
		baseUrl := fmt.Sprintf("%s://%s", urlObj.Scheme, urlObj.Host)
		for {
			select {
			case <-ticker.C:
				log.Println("ticker running ...")
				location := ui.Eval("window.location.href")
				log.Printf("location :%s\n", location)
				href := location.String()
				if strings.HasPrefix(href, baseUrl) {
					if strings.HasSuffix(href, "index.cgi") {
						rawCookies := ui.Eval("document.cookie")
						log.Printf("rawCookies :%s\n", rawCookies)
						header := http.Header{}
						header.Add("Cookie", rawCookies.String())
						request := http.Request{Header: header}
						cookieToken := ""
						for _, cookie := range request.Cookies() {
							if cookie.Name == "DSID" {
								cookieToken = cookie.Value
							}
						}
						if cookieToken != "" {
							cookieInterval.cookie = cookieToken
							// ui.Close()
							stop(cookieInterval)
						}
					} else if strings.HasSuffix(href, "welcome.cgi") {
						ui.Eval("window['frmLogin']['sn-preauth-proceed'].click();")
					} else if strings.HasSuffix(href, "confirm") {
						ui.Eval("window['frmConfirmation']['btnContinue'].click();")
					}
				} else if strings.HasPrefix(href, "https://login.microsoftonline.com") {
					ui.Eval(fmt.Sprintf("el=document.querySelector(\"input[type='email']\");el.value = '%s';el.blur();document.querySelector(\"input[type='submit']\").click()", email))
				} else if strings.HasPrefix(href, "https://sts.") {
					ui.Eval(fmt.Sprintf("el=document.querySelector(\"input[type='password']\");el.value = '%s';el.blur();document.querySelector('form').submit();", password))
				}
			case <-quit:
				ticker.Stop()
				log.Println("ticker stopped")
				return
			}
		}
	}()
	return cookieInterval
}

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
	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	// launch Chrome Instance (lorca)
	ui, err := lorca.New("", "", 1024, 768, args...)
	log.Println(("hello 2"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(("hello 3"))
	ui.Bind("login", login)
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
	// url := "http://localhost:5000"
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
