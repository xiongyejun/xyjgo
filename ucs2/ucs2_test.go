package ucs2

import (
	"io/ioutil"
	"testing"
)

func Test_func1(t *testing.T) {
	b, _ := ioutil.ReadFile("utf8.txt")

	if b, err := FromUTF8(b); err != nil {
		t.Log(err)
	} else {
		if err := ioutil.WriteFile("ucs2.txt", b, 0666); err != nil {
			t.Log(err)
		}
	}

}

// func Test_func2(t *testing.T) {
// 	if b, err := ioutil.ReadFile("ucs2.txt"); err != nil {
// 		t.Log(err)
// 	} else {
// 		if buf8, err := ToUTF8(b); err != nil {
// 			t.Log(err)
// 		} else {
// 			t.Logf("%s\n", buf8)
// 		}
// 	}
// }
