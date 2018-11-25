package translate

type BaiDu struct {
	tsl
}

func NewBaiDu() *BaiDu {
	b := new(BaiDu)
	b.url = "https://fanyi.baidu.com/"
	return b
}

func (me *BaiDu) Translate(value string) (ret string, err error) {
	return
}
