package OCRService

import (
	"beianAPI/utils"
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func TrWebOCRService(base64 string) (string, error) {
	config := &utils.Conf.TROCR
	OCRUrl := config.Url
	res, err := http.PostForm(OCRUrl, url.Values{
		"img": {base64},
	})

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	re := regexp.MustCompile("\"[0-9]{0,4}\"")
	if matchStrs := re.FindStringSubmatch(buf.String()); len(matchStrs) > 0 {
		result := strings.Trim(matchStrs[0], "\"")
		return result, err
	}
	return "", err
}
