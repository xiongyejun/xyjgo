package translate

import (
	"encoding/json"
)

type YouDao struct {
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

func NewYouDao() *YouDao {
	y := new(YouDao)
	y.url = "http://fanyi.youdao.com/translate"
	return y
}

func (me *YouDao) Translate(value string) (ret string, err error) {
	var b []byte
	if b, err = httpPost(me.url, "i="+value+"&doctype=json"); err != nil {
		return
	}

	return me.getResult(b)
}

func (me *YouDao) getResult(b []byte) (ret string, err error) {
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
		}
	}

	return
}
