package sougou

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type SouGou struct {
	URL string
}

type s1 struct {
	Content string   `json:"content"`
	Frame   []string `json:"frame"`
}

type s0 struct {
	Result  []s1    `json:"result"`
	Success float64 `json:"success"`
}

func init() {

}

func (me *SouGou) OCR(picPath string) (ret string, err error) {
	var b []byte
	// 上传图片
	if b, err = httpPost("http://pic.sogou.com/pic/upload_pic.jsp", "pic_path="+picPath); err != nil {
		return
	}
	return string(b), nil
}

func New() (ret *SouGou) {
	ret = new(SouGou)
	ret.URL = "http://ocr.shouji.sogou.com/v2/ocr/json"
	return
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
