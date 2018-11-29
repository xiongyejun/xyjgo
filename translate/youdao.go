package translate

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/xiongyejun/xyjgo/winAPI/playMP3"
)

type youDao struct {
	tsl
}

type s1 struct {
	Src string `json:"Src"`
	Tgt string `json:"Tgt"`
}
type s0 struct {
	TranslateResult [][]s1  `json:"TranslateResult"`
	Type            string  `json:"Type"`
	ErrorCode       float64 `json:"ErrorCode"`
	ElapsedTime     float64 `json:"ElapsedTime"`
}

func NewYouDao() (y *youDao, err error) {
	y = new(youDao)
	y.url = "http://fanyi.youdao.com/translate"
	//	y.mp3URL = "http://dict.youdao.com/dictvoice?audio=" // + value + &type=2
	y.mp3URL = "http://tts.youdao.com/fanyivoice?" // + &le=eng&keyfrom=speaker-target
	y.mp3SavePath = os.Getenv("USERPROFILE") + `\YouDaoMP3\`
	// 检查文件夹是否存在
	if _, err = os.Stat(y.mp3SavePath); err != nil {
		if err = os.Mkdir(y.mp3SavePath, 0666); err != nil {
			return
		}
	}

	return
}

// 翻译
func (me *youDao) Translate(value string, bSpeak bool) (ret string, err error) {
	var b []byte
	var tgt string
	if b, err = httpPost(me.url, "i="+value+"&doctype=json"); err != nil {
		return
	}

	if ret, tgt, err = me.getResult(b); err != nil {
		return
	}

	if bSpeak {
		var speakValue string = tgt
		if (value[0] >= 'a' && value[0] <= 'z') || (value[0] >= 'A' && value[0] <= 'Z') {
			speakValue = value
		}
		if err = me.speak(speakValue); err != nil {
			fmt.Println(errors.New("speak出错：" + err.Error()))
		}
	}

	return ret, nil
}

// 朗读
func (me *youDao) speak(value string) (err error) {
	// 名称都用sha1保存
	sha := sha1.New()
	if _, err = sha.Write([]byte(value)); err != nil {
		return
	}
	strSha := hex.EncodeToString(sha.Sum(nil)) + ".mp3"
	filePath := me.mp3SavePath + strSha
	// 判断是否已经存在
	if _, err = os.Stat(filePath); err != nil {
		// 不存在就下载
		var bMP3 []byte
		if bMP3, err = me.getMP3(value); err != nil {
			return
		}

		// 保存MP3
		if err = ioutil.WriteFile(filePath, bMP3, 0666); err != nil {
			return
		}
	}
	// 朗读MP3
	if err = playMP3.Play(filePath, true); err != nil {
		return
	}
	return nil
}

// 下载MP3
func (me *youDao) getMP3(value string) (bMP3 []byte, err error) {
	u := url.Values{}
	u.Set("word", value)
	return httpGet(me.mp3URL + u.Encode() + "&le=eng&keyfrom=speaker-target")
}

// 获取翻译需要的信息
// 需要的结果在TranslateResult里面
// ret	返回的是翻译前和翻译后在一起的string
// tgt	仅仅是翻译后的
func (me *youDao) getResult(b []byte) (ret string, tgt string, err error) {
	s := new(s0)

	if err = json.Unmarshal(b, &s); err != nil {
		return
	}

	for i := range s.TranslateResult {
		for j := range s.TranslateResult[i] {
			ret += s.TranslateResult[i][j].Src
			ret += ":"
			ret += s.TranslateResult[i][j].Tgt
			ret += "\r\n"

			tgt += s.TranslateResult[i][j].Tgt
		}
	}

	return
}
