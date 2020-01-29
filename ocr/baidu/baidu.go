// https://ai.baidu.com/docs#/OCR-API-AccurateBasic/top
package baidu

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type token struct {
	Access_token   string  `json:"access_token"`
	Scope          string  `json:"scope"`
	Session_secret string  `json:"session_secret"`
	Refresh_token  string  `json:"refresh_token"`
	Expires_in     float64 `json:"expires_in"`
	Session_key    string  `json:"session_key"`
}

type BaiDu struct {
	URL string
	token
}

type result struct {
	Words string `json:"words"`
}

type results struct {
	Words_result     []result `json:"words_result"`
	Log_id           int      `json:"log_id"`
	Words_result_num int      `json:"words_result_num"`
}

func (me *BaiDu) OCR(picPath string) (ret string, err error) {
	/*
		参数						是否必选	类型	可选值范围	说明
		image						true		string	-			图像数据，base64编码后进行urlencode，要求base64编码和urlencode后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/jpeg/png/bmp格式
		language_type				false		string				识别语言类型，默认为CHN_ENG
		detect_direction			false		string	true、false	是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括:
																	- true：检测朝向；
																	- false：不检测朝向。
		paragraph					false		string	true/false	是否输出段落信息
		probability					false		string	true/false	是否返回识别结果中每一行的置信度
	*/

	var b []byte
	if b, err = ioutil.ReadFile(picPath); err != nil {
		return
	}

	picstr := base64.StdEncoding.EncodeToString(b)
	picstr = url.QueryEscape(picstr)

	strPost := "image=" + picstr
	if b, err = httpPost(me.URL+"?access_token="+me.Access_token, strPost); err != nil {
		return
	}
	var rs results = results{}
	if err = json.Unmarshal(b, &rs); err != nil {
		return
	}
	var ss []string = make([]string, rs.Words_result_num)
	for i := range rs.Words_result {
		ss[i] = rs.Words_result[i].Words
	}
	ret = strings.Join(ss, "\n")

	return
}

func New() (ret *BaiDu, err error) {
	ret = new(BaiDu)
	ret.URL = "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate_basic"

	if err = ret.get_access_token(); err != nil {
		return
	}

	return
}

// access_token	通过API Key和Secret Key获取的access_token
func (me *BaiDu) get_access_token() (err error) {
	/*
		   获取Access Token
		   请求URL数据格式

		   向授权服务地址https://aip.baidubce.com/oauth/2.0/token发送请求（推荐使用POST），并在URL中带上以下参数：

		   grant_type： 必须参数，固定为client_credentials；
		   client_id： 必须参数，应用的API Key；
		   client_secret： 必须参数，应用的Secret Key；

			服务器返回的JSON文本参数如下：

			access_token： 要获取的Access Token；
			expires_in： Access Token的有效期(秒为单位，一般为1个月)；
			其他参数忽略，暂时不用;

	*/

	/*
	   若请求错误，服务器将返回的JSON文本包含以下参数：

	   error： 错误码；关于错误码的详细信息请参考下方鉴权认证错误码。
	   error_description： 错误描述信息，帮助理解和解决发生的错误。
	*/
	var client_id string = "E4AEhxW2bXR0P6qFlVMQ2GtR"
	var client_secret string = "oxBVuCaeCANY6rZF9CaGlz6zhvIC3bFg"

	strPost := "grant_type=client_credentials&client_id=" + client_id + "&client_secret=" + client_secret
	var b []byte
	if b, err = httpPost("https://aip.baidubce.com/oauth/2.0/token", strPost); err != nil {
		return
	}

	me.token = token{}
	if err = json.Unmarshal(b, &me.token); err != nil {
		return
	}

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
