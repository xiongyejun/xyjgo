package google

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/url"

	"github.com/xiongyejun/xyjgo/translate"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type google struct {
	translate.Tsl
}

func New() (ret *google, err error) {
	ret = new(google)
	ret.URL = "https://translate.google.cn/m?"

	return
}

func (me *google) Translate(value string) (ret string, err error) {
	var strTL string
	if (value[0] >= 'a' && value[0] <= 'z') || (value[0] >= 'A' && value[0] <= 'Z') {
		strTL = `&hl=zh-CN&sl=en&tl=zh-CN`
	} else {
		strTL = `&hl=en&sl=zh-CN&tl=en`
	}

	value = url.QueryEscape(value)
	var b []byte
	// https://translate.google.cn/m?hl=zh-CN&sl=en&tl=zh-CN&ie=UTF-8&prev=_m&q=
	if b, err = translate.HttpGet(me.URL + strTL + `&ie=UTF-8&prev=_m&q=` + value); err != nil {
		return
	}

	// <input type="text" aria-label="Source text" name="q" class="input-field" maxlength="2048" value="密码"><div class="translate-button-container"><input type="submit" value="Translate" class="translate-button"></div></form></div><div class="result-container">password</div>
	b = bytes.Split(b, []byte(`<div class="result-container">`))[1] // <div dir="ltr" class="t0">你好吗</div>
	b = bytes.Split(b, []byte(`</div>`))[0]
	if b, err = gbkToUtf8(b); err != nil {
		return
	}
	b = bytes.ReplaceAll(b, []byte(`&#39;`), []byte("'"))
	b = bytes.ReplaceAll(b, []byte(`&#160;`), []byte(" "))
	ret = string(b)

	return
}

func (me *google) Speak(value string) (err error) {
	return errors.New("未实现")
}

func gbkToUtf8(b []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(b), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
