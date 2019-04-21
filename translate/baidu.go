// http://api.fanyi.baidu.com/api/trans/product/apidoc
package translate

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
)

type baiDu struct {
	tsl

	q     string
	from  string
	to    string
	appid string
	salt  string
	sign  string // appid+q+salt+密钥 的MD5值
	key   string
}

type BaiduErrS struct {
	Error_msg  string `json:"error_msg"`
	Error_code string `json:"error_code"`
}
type BaiduS1 struct {
	Dst string `json:"dst"`
	Src string `json:"src"`
}

type BaiduS0 struct {
	From         string    `json:"from"`
	To           string    `json:"to"`
	Trans_result []BaiduS1 `json:"trans_result"`
}

func NewBaiDu() (b *baiDu, err error) {
	b = new(baiDu)
	b.url = "http://api.fanyi.baidu.com/api/trans/vip/translate"
	b.appid = "20190421000290122"
	b.key = "OmV8zMcBZTSHzztWrLNY"

	return
}

func (me *baiDu) Translate(value string, bSpeak bool) (ret string, err error) {
	var sPost []string = make([]string, 0)

	me.q = value
	sPost = append(sPost, "q="+url.QueryEscape(me.q))
	if (value[0] >= 'a' && value[0] <= 'z') || (value[0] >= 'A' && value[0] <= 'Z') {
		sPost = append(sPost, "from=en")
		sPost = append(sPost, "to=zh")
	} else {
		sPost = append(sPost, "from=zh")
		sPost = append(sPost, "to=en")
	}
	sPost = append(sPost, "appid="+me.appid)

	me.getSalt()
	if err = me.getSign(); err != nil {
		return
	}
	sPost = append(sPost, "salt="+me.salt)
	sPost = append(sPost, "sign="+me.sign)

	strPost := strings.Join(sPost, "&")

	var b []byte
	if b, err = httpPost(me.url, strPost); err != nil {
		return
	}

	return me.getResult(b)
}

func (me *baiDu) getResult(b []byte) (ret string, err error) {
	s := new(BaiduS0)

	if err = json.Unmarshal(b, &s); err != nil {
		return
	}

	if len(s.Trans_result) > 0 {
		for i := range s.Trans_result {
			ret += s.Trans_result[i].Dst
			ret += "\r\n"
		}
	} else {
		sErr := new(BaiduErrS)

		if err = json.Unmarshal(b, &sErr); err != nil {
			return
		}

		var ok bool
		if ret, ok = baiduErr[sErr.Error_code]; ok {
			return
		} else {
			return "error_code:" + sErr.Error_code, nil
		}
	}

	return
}

func (me *baiDu) getSalt() {
	n := rand.Int()
	me.salt = strconv.Itoa(n)

	return
}
func (me *baiDu) getSign() (err error) {
	// 在生成签名拼接 appid+q+salt+密钥 字符串时，q不需要做URL encode，在生成签名之后，发送HTTP请求之前才需要对要发送的待翻译文本字段q做URL encode。
	str := me.appid + me.q + me.salt + me.key
	hash := md5.New()
	if _, err = hash.Write([]byte(str)); err != nil {
		return
	}
	b := hash.Sum(nil)
	me.sign = hex.EncodeToString(b)

	return
}
func (me *baiDu) speak(value string) (err error) {
	return errors.New("未实现")
}
