package crawler

import (
	"beianAPI/utils/OCRService"
	"encoding/base64"
	"github.com/gocolly/colly"
	"math/rand"
	"net/http"
)

const (
	letterBytes       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	primaryDomain     = "beian.gov.cn"
	secondLevelDomain = "www.beian.gov.cn"
	recordQueryUrl    = "http://www.beian.gov.cn/portal/recordQuery"
	urlCAPTCHA        = "http://www.beian.gov.cn/common/image.jsp?t=2"
)

func randomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func traverseBase64(img []byte) string {
	res := base64.StdEncoding.EncodeToString(img)
	return res
}

func getTokenAndCookie(domainName string) (map[string]string, error) {

	c := colly.NewCollector(
		colly.AllowedDomains(primaryDomain, secondLevelDomain),
		colly.AllowURLRevisit(),
	)

	userAgent := randomString()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", userAgent)
		r.Headers.Set("Upgrade-Insecure-Requests", "1")
	})

	var cookie []*http.Cookie
	c.OnResponse(func(r *colly.Response) {
		cookie = c.Cookies(recordQueryUrl)
	})

	if err := c.Visit(recordQueryUrl); err != nil {
		return map[string]string{}, err
	}

	if err := c.SetCookies("*", cookie); err != nil {
		return map[string]string{}, err
	}

	var token string
	c.OnHTML("a:contains(查询)", func(e *colly.HTMLElement) {
		if e.Text == "查询" {
			context, _ := e.DOM.Attr("onclick")
			token = context[48 : len(context)-1]
		}
	})

	if err := c.Visit(recordQueryUrl); err != nil {
		return map[string]string{}, err
	}

	if err := c.SetCookies(urlCAPTCHA, cookie); err != nil {
		return map[string]string{}, err
	}
	imageC := c.Clone()

	var CAPTCHA []byte
	imageC.OnResponse(func(r *colly.Response) {
		CAPTCHA = r.Body
	})

	if err := imageC.Visit(urlCAPTCHA); err != nil {
		return map[string]string{}, err
	}

	passWord, err := OCRService.TrWebOCRService(traverseBase64(CAPTCHA))
	if err != nil {
		return map[string]string{}, err
	}

	//c.OnResponse(func(r *colly.Response) {
	//	f, err := os.Create("resp.html")
	//	if err != nil {
	//		panic(err)
	//	}
	//	io.Copy(f, bytes.NewReader(r.Body))
	//})

	cookies := cookie[0].Name + "=" + cookie[0].Value + "; " + cookie[1].Name + cookie[1].Value
	authInfos := map[string]string{
		authToken:     token,
		authPassWord:  passWord,
		authUserAgent: userAgent,
		authCookie:    cookies,
	}

	return authInfos, err
}
