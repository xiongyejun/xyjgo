package baidu

import (
	"testing"
)

func Test_func(t *testing.T) {
	var bd *BaiDu
	var err error
	if bd, err = New(); err != nil {
		t.Log(err)
		return
	}
	t.Log(bd.Access_token)

	var str string
	if str, err = bd.OCR(`..\test.png`); err != nil {
		t.Log(err)
		return
	}
	t.Log(str)
}
