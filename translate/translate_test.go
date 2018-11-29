package translate

import (
	"testing"
)

func Test_func(t *testing.T) {
	var b ITranslate
	var err error
	if b, err = NewYouDao(); err != nil {
		t.Error(err)
	}

	var ret string
	var tgt string
	if ret, tgt, err = b.Translate("你好，我是谁？你看过射雕英雄传吗？"); err != nil {
		t.Error(err)
	} else {
		print(ret)
	}

	if err = b.Speak(tgt); err != nil {
		t.Error(err)
	}
}
