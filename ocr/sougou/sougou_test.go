package sougou

import (
	"testing"
)

func Test_f(t *testing.T) {
	sg := New()
	ret, err := sg.ConvPic(`C:\Users\Administrator\Desktop\微信图片_20191015191957.jpg`)
	t.Log(ret)
	t.Log(err)
}
