package crawler

import (
	"beianAPI/model"
	"io"
	"runtime"
	"sync"
)

const (
	maxTries     = 100
	maxGoRoutine = 3

	authToken     = "token"
	authCookie    = "cookies"
	authPassWord  = "passWord"
	authUserAgent = "userAgent"

	siteName   = "网站名称"
	siteDomain = "网站主域名"
	siteEntity = "开办主体"
	siteClass  = "网站类别"

	operatorName = "开办者名称"
	recordICP    = "公安备案号"
	recordeDept  = "备案地公安机关 "
	recordTime   = "联网备案时间"
)

func GetBeiAnInfo(domainName string) *model.Beian {
	if infoMap := publishTask(domainName); len(infoMap) > 0 {
		return &model.Beian{
			SiteName:   infoMap[siteName],
			SiteDomain: infoMap[siteDomain],
			SiteEntity: infoMap[siteEntity],
			SiteClass:  infoMap[siteClass],

			OperatorName: infoMap[operatorName],
			RecordICP:    infoMap[recordICP],
			RecordeDept:  infoMap[recordeDept],
			RecordTime:   infoMap[recordTime],
		}
	}
	return nil
}

func publishTask(domainName string) map[string]string {
	infoMapChan := make(chan map[string]string, maxTries)
	var wg sync.WaitGroup
	wg.Add(maxTries)

	// limit number of crawlers
	runtime.GOMAXPROCS(maxGoRoutine)
	for i := 0; i < maxTries; i++ {
		go crawInfo(domainName, infoMapChan, &wg)
	}

	wg.Wait()

	if len(infoMapChan) > 0 {
		return <-infoMapChan
	}
	return map[string]string{}
}

func crawInfo(domainName string, infoMapChan chan<- map[string]string, wg *sync.WaitGroup) {
	defer wg.Done()
	var authInfo map[string]string
	var err error
	var infoMap map[string]string
	if authInfo, err = getTokenAndCookie(domainName); err != nil {
		return
	}

	var body io.Reader
	if body, err = getHTML(domainName, authInfo[authToken], authInfo[authPassWord],
		authInfo[authUserAgent], authInfo[authCookie]); err != nil {
		return
	}

	if infoMap, err = requestJudge(body); err != nil {
		return
	}

	if len(infoMap) > 0 {
		infoMapChan <- infoMap
	}

}
