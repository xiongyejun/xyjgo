package translate

import (
	"testing"
)

func Test_func(t *testing.T) {
	var b ITranslate
	var err error
	if b, err = NewBaiDu(); err != nil {
		t.Error(err)
	}

	var ret string
	if ret, err = b.Translate("Have you seen the legend of the condor heroes?", false); err != nil {
		t.Error(err)
		return
	} else {
		print(ret)
	}

	if ret, err = b.Translate("你是谁？你好吗？", true); err != nil {
		t.Error(err)
	} else {
		print(ret)
	}
}
