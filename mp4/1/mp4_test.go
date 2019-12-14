package mp4

import (
	"os"
	"testing"
)

func Test_func(t *testing.T) {
	var f *os.File
	var err error
	if f, err = os.Open("../1.mp4"); err != nil {
		t.Error(err)
		return
	}
	defer f.Close()

	m, err := Decode(f)

	ft := Clip(5*60, 30)
	var fw *os.File
	if fw, err = os.OpenFile("2.mp4", os.O_CREATE|os.O_WRONLY, 0666); err != nil {
		t.Error(err)
		return
	}

	defer fw.Close()

	EncodeFiltered(fw, m, ft)
}
