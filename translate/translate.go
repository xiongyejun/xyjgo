// 翻译

package translate

import (
	"io/ioutil"
	"net/http"
	"strings"
)

var baiduErr map[string]string = map[string]string{
	"52000": "成功",
	"52001": "请求超时",
	"52002": "系统错误",
	"52003": "未授权用户",
	"54000": "必填参数为空",
	"54001": "签名错误",
	"54003": "访问频率受限",
	"54004": "账户余额不足",
	"54005": "长query请求频繁",
	"58000": "客户端IP非法",
	"58001": "译文语言方向不支持",
	"58002": "服务当前已关闭",
}

type ITranslate interface {
	Translate(value string, bSpeak bool) (ret string, err error)
	speak(value string) (err error)
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
