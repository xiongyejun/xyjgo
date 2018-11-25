// 翻译

package translate

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type ITranslate interface {
	Translate(value string) (ret string, err error)
}

type tsl struct {
	url string
}

func httpPost(url string, strPost string) (ret []byte, err error) {
	var resp *http.Response
	if resp, err = http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(strPost)); err != nil {
		return
	}
	defer resp.Body.Close()
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	return body, err
}
