package DATA

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type Webinfo struct {
	Url      string
	Title    string
	Status   string
	Length   string
	Redirect string
	Server   string
}

var Webinfos []Webinfo

func (webinfo *Webinfo) GETWEBINFO(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, respbody, _, err := RequestHead("GET", url, nil, map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.72 Safari/537.36", "Cookie": "rememberMe=12"})
	if err != nil {
		return
	}
	webinfo.Length, webinfo.Title = GETTitleLength(respbody)
	for _, rule := range RuleDatas {
		switch rule.Type {
		case "code":
			Match, _ := regexp.MatchString(rule.Rule, respbody)
			if Match {
				webinfo.Server = rule.Name
			}
		case "headers":
			Match, _ := regexp.MatchString(rule.Rule, fmt.Sprintf("%s", resp.Header))
			if Match {
				webinfo.Server = rule.Name
			}
		case "title":
			Match, _ := regexp.MatchString(rule.Rule, webinfo.Title)
			if Match {
				webinfo.Server = rule.Name
			}
		default:
			webinfo.Server = resp.Header.Get("Server")
		}

	}
	webinfo.Status = resp.Status
	webinfo.Url = url
	fmt.Println(webinfo)
}
func GETTitleLength(respbody string) (string, string) {
	html := strings.ToLower(respbody)
	regex1, _ := regexp.Compile("<title>(.*?)</title>")
	titles := regex1.FindStringSubmatch(html)
	var Titel, Rlength string
	if len(titles) < 1 {
		Titel = ""
	} else {
		Titel = titles[1]
	}
	if len(respbody) >= 1000 {
		Rlength = strconv.Itoa(len(respbody)/1000) + "KB"
	} else {
		Rlength = strconv.Itoa(len(respbody)) + "B"
	}
	return Rlength, Titel

}
