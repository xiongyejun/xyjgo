package sougou

import (
	"testing"
)

func Test_f(t *testing.T) {
	sg := New()
	ret, err := sg.ConvPic(`C:\Users\Administrator\Desktop\1.png`)
	t.Log(ret)
	t.Log(err)
}
