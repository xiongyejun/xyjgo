// 翻译

package translate

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type ITranslate interface {
	Translate(value string) (ret string, tgt string, err error)
	Speak(value string) (err error)
}

type tsl struct {
	url         string
	mp3URL      string
	mp3SavePath string
}

func httpPost(url string, strPost string) (ret []byte, err error) {
	var resp *http.Response
	if resp, err = http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(strPost)); err != nil {
		return
	}
	defer resp.Body.Close()
	if ret, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	return
}

func httpGet(url string) (ret []byte, err error) {
	var resp *http.Response
	if resp, err = http.Get(url); err != nil {
		return
	}
	defer resp.Body.Close()
	if ret, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	return
}
