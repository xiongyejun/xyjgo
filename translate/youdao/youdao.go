package youdao

import (
	"errors"
	"runtime"

	"encoding/json"
	"net/url"
	"os"

	"github.com/xiongyejun/xyjgo/translate"
)

type youDao struct {
	translate.Tsl
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

func New() (y *youDao, err error) {
	y = new(youDao)
	y.URL = "http://fanyi.youdao.com/translate"
	//	y.mp3URL = "http://dict.youdao.com/dictvoice?audio=" // + value + &type=2
	y.Mp3URL = "http://tts.youdao.com/fanyivoice?" // + &le=eng&keyfrom=speaker-target

	switch runtime.GOOS {
	case "darwin":
		y.Mp3SavePath = os.Getenv("HOME") + `/YouDaoMP3/`
	case "windows":
		y.Mp3SavePath = os.Getenv("USERPROFILE") + `\YouDaoMP3\`
	default:
		return nil, errors.New("未知系统.")
	}

	// 检查文件夹是否存在
	if _, err = os.Stat(y.Mp3SavePath); err != nil {
		if err = os.Mkdir(y.Mp3SavePath, 0666); err != nil {
			return
		}
	}

	return
}

// 翻译
func (me *youDao) Translate(value string) (ret string, err error) {
	var b []byte
	// var tgt string
	if b, err = translate.HttpPost(me.URL, "i="+value+"&doctype=json"); err != nil {
		return
	}

	if ret, _, err = me.getResult(b); err != nil {
		return
	}

	return ret, nil
}

// // 朗读
func (me *youDao) Speak(value string) (err error) {
	// 	// 名称都用sha1保存
	// 	sha := sha1.New()
	// 	if _, err = sha.Write([]byte(value)); err != nil {
	// 		return
	// 	}
	// 	strSha := hex.EncodeToString(sha.Sum(nil)) + ".mp3"
	// 	filePath := me.mp3SavePath + strSha
	// 	// 判断是否已经存在
	// 	if _, err = os.Stat(filePath); err != nil {
	// 		// 不存在就下载
	// 		var bMP3 []byte
	// 		if bMP3, err = me.getMP3(value); err != nil {
	// 			return
	// 		}

	// 		// 保存MP3
	// 		if err = ioutil.WriteFile(filePath, bMP3, 0666); err != nil {
	// 			return
	// 		}
	// 	}
	// 	// 朗读MP3
	// 	if err = playMP3.Play(filePath, true); err != nil {
	// 		return
	// 	}
	return nil
}

// 下载MP3
func (me *youDao) getMP3(value string) (bMP3 []byte, err error) {
	u := url.Values{}
	u.Set("word", value)
	return translate.HttpGet(me.Mp3URL + u.Encode() + "&le=eng&keyfrom=speaker-target")
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
