package rleVBA

import (
	"io/ioutil"
	"testing"
)

func Test_func(t *testing.T) {
	b, _ := ioutil.ReadFile(`C:\Users\Administrator\Documents\myvba\VBAProject\testdata\vbaProject\VBA\dir`)

	r := NewRLE(b)

	b = r.UnCompress()

	ioutil.WriteFile("1.txt", b, 0666)
}
