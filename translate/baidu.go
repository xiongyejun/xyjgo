package translate

import (
	"errors"
)

type baiDu struct {
	tsl
}

func NewBaiDu() (b *baiDu, err error) {
	b = new(baiDu)
	b.url = "https://fanyi.baidu.com/"
	return nil, errors.New("未实现")
}

func (me *baiDu) Translate(value string, bSpeak bool) (ret string, err error) {
	return "", errors.New("未实现")
}
func (me *baiDu) speak(value string) (err error) {
	return errors.New("未实现")
}