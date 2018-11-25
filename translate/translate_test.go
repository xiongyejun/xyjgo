package translate

import (
	"testing"
)

func Test_func(t *testing.T) {
	var b ITranslate = NewBaiDu()
	if ret, err := b.Translate("ceshi"); err != nil {
		t.Error(err)
	} else {
		t.Log(ret)
	}

	b = NewYouDao()
	if ret, err := b.Translate("你好，我是谁？你看过射雕英雄传吗？"); err != nil {
		t.Error(err)
	} else {
		print(ret)
	}
}
