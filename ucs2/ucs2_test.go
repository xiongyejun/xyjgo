package ucs2

import (
	"io/ioutil"
	"testing"
)

func Test_func(t *testing.T) {
	//	var str string = "ucs2T0utf8测试一下吧"

	//	if b, err := FromUTF8([]byte(str)); err != nil {
	//		t.Log(err)
	//	} else {
	//		if err := ioutil.WriteFile("ucs2.txt", b, 0666); err != nil {
	//			t.Log(err)
	//		}
	//	}

	if err := testfunc2(); err != nil {
		t.Log(err)
	}
}

func testfunc2() error {
	if b, err := ioutil.ReadFile("ucs2.txt"); err != nil {
		return err
	} else {
		if buf8, err := ToUTF8(b); err != nil {
			return err
		} else {
			print(string(buf8))
		}
	}
	return nil
}
