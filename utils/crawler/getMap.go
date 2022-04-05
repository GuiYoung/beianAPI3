package crawler

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	registerSystemInfoURL = "http://www.beian.gov.cn/portal/registerSystemInfo"
)

func getHTML(domainName, token, pwd, useragent, cookie string) (io.Reader, error) {

	method := "POST"

	urlForm := "token=" + token + "&sdcx=1&flag=2&domainname=" + domainName + "&inputPassword=" + pwd
	payload := strings.NewReader(urlForm)

	client := &http.Client{}
	req, err := http.NewRequest(method, registerSystemInfoURL, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Cookie", cookie)
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", useragent)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	//f, err := os.Create("resp.html")
	//if err != nil {
	//	return nil,err
	//}
	//
	//io.Copy(f, bytes.NewReader(body))

	return bytes.NewReader(body), nil
}

// judge is get correct html code, correctly get, then return map with info
func requestJudge(body io.Reader) (infoMap map[string]string, err error) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return infoMap, err
	}

	lastField := ""
	fDoc := doc.Find(".wzjb")
	infoMap = map[string]string{}

	//fmt.Println(fDoc.Size())

	if fDoc.Size() != 0 {
		fDoc.Children().Each(func(i int, s *goquery.Selection) {
			s.Find("td").Each(func(i int, s *goquery.Selection) {
				infoMap[lastField] = s.Text()
				lastField = s.Text()
			})
		})
	}
	return
}
