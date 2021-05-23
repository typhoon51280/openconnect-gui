package lorca

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/zserge/lorca"
)

type CookieInterval struct {
	running chan bool
	cookie  string
}

func Login(loginUrl string, email string, password string, headless bool) string {
	log.Printf("login %s\n", loginUrl)
	args := []string{}
	if headless {
		args = append(args, "--headless")
	}
	ui, err := lorca.New(loginUrl, "", 1000, 600, args...)
	if err != nil {
		log.Fatal(err)
	}
	ui.Load(loginUrl)
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
					homePage := fmt.Sprintf("%s%s", baseUrl, "/dana/home/")
					if strings.HasPrefix(href, homePage) {
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
