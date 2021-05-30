package lorca

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/koltyakov/lorca"
)

type CookieMonitor struct {
	url        url.URL
	cookies    *Cookies
	tokenName  string
	tokenValue string
	isRunning  bool
	state      chan bool
}

var cookieCache *Cookies = &Cookies{}

func Login(loginUrl string, email string, password string, headless bool) string {
	log.Printf("login %s\n", loginUrl)
	args := []string{}
	if headless {
		args = append(args, "--headless")
	}
	// args = append(args, "--proxy-server=http://localhost:9999")
	loginWindow, err := lorca.New("", "", 1000, 600, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer loginWindow.Close()
	log.Printf("cookieCache: %s", cookieCache.toString())
	setCookies(loginWindow, cookieCache)
	loginWindow.Load(loginUrl)
	log.Println("login opened")
	token := openWindow(loginWindow, loginUrl, "DSID", email, password)
	log.Println("login closed")
	return token
}

func (cookieMonitor *CookieMonitor) wait() {
	log.Println("CookieMonitor Wait")
	cookieMonitor.isRunning = true
	<-cookieMonitor.state
}

func (cookieMonitor *CookieMonitor) toggle(state bool) {
	cookieMonitor.isRunning = state
	cookieMonitor.state <- cookieMonitor.isRunning
}

func (cookieMonitor *CookieMonitor) stop() {
	if cookieMonitor.isRunning {
		cookieMonitor.toggle(false)
	}
	log.Println("CookieMonitor Stop")
}

func (cookieMonitor *CookieMonitor) close() {
	close(cookieMonitor.state)
	log.Println("CookieMonitor Close")
}

func (cookieMonitor *CookieMonitor) save(window lorca.UI) {
	cookieMonitor.cookies = getCookies(window)
	cookieCache = cookieMonitor.cookies
	cookieMonitor.tokenValue = cookieMonitor.findCookie(cookieMonitor.tokenName)
	log.Printf("CookieMonitor Saved: %s", cookieCache.toString())
}

func (cookieMonitor *CookieMonitor) findCookie(name string) string {
	cookie := cookieMonitor.cookies.find(cookieMonitor.url.Hostname(), name)
	if cookie != nil {
		return cookie.Value
	}
	return ""
}

func getCookies(window lorca.UI) *Cookies {
	cookies := &Cookies{}
	response := window.Send("Network.getCookies", nil)
	if response.Err() != nil {
		log.Printf("%s", response.Err())
		return cookies
	}
	if err := response.Object()["cookies"].To(cookies); err != nil {
		log.Printf("%s", err)
		return cookies
	}
	return cookies
}

func setCookies(window lorca.UI, cookies *Cookies) {
	response := window.Send("Network.setCookies", cookies.toMap())
	if response.Err() != nil {
		log.Printf("%s", response.Err())
	}
}

func openWindow(window lorca.UI, loginUrl string, cookieName string, email string, password string) string {
	ticker := time.NewTicker(2 * time.Second)
	urlObj, err := url.Parse(loginUrl)
	if err != nil {
		log.Fatal(err)
	}
	cookieMonitor := &CookieMonitor{
		isRunning:  false,
		state:      make(chan bool),
		url:        *urlObj,
		cookies:    cookieCache,
		tokenName:  cookieName,
		tokenValue: "",
	}
	go func() {
		domainURL := fmt.Sprintf("%s://%s", urlObj.Scheme, urlObj.Host)
		for {
			select {
			case <-ticker.C:
				log.Println("ticker running ...")
				if cookieMonitor.isRunning {
					location := window.Eval("window.location.href")
					if location.Err() != nil {
						log.Printf("%s", location.Err())
						return
					}
					href := location.String()
					log.Printf("location :%s\n", href)
					docCookies := window.Eval("document.cookie")
					if docCookies.Err() == nil {
						log.Printf("docCookies: %s", docCookies.String())
					}
					if strings.HasPrefix(href, domainURL) {
						homePage := fmt.Sprintf("%s%s", domainURL, "/dana/home/")
						if strings.HasPrefix(href, homePage) {
							cookieMonitor.save(window)
							cookieMonitor.stop()
						} else if strings.HasSuffix(href, "welcome.cgi") {
							window.Eval("window['frmLogin']['sn-preauth-proceed'].click();")
						} else if strings.HasSuffix(href, "confirm") {
							window.Eval("window['frmConfirmation']['btnContinue'].click();")
						}
					} else if strings.HasPrefix(href, "https://login.microsoftonline.com") {
						window.Eval(fmt.Sprintf("el=document.querySelector(\"input[type='email']\");el.value = '%s';el.blur();document.querySelector(\"input[type='submit']\").click()", email))
					} else if strings.HasPrefix(href, "https://sts.") {
						window.Eval(fmt.Sprintf("el=document.querySelector(\"input[type='password']\");el.value = '%s';el.blur();document.querySelector('form').submit();", password))
					}
				}
			case <-cookieMonitor.state:
				log.Println("cookieMonitor.state ...")
				if !cookieMonitor.isRunning {
					ticker.Stop()
					log.Println("cookieMonitor.state, ticker stopped")
					return
				}
			case <-window.Done():
				if cookieMonitor.isRunning {
					log.Println("window.Done")
					cookieMonitor.stop()
				}
			}
		}
	}()
	cookieMonitor.wait()
	cookieMonitor.close()
	return cookieMonitor.tokenValue
}
