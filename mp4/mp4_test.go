package mp4

import (
	"os"
	"testing"
	"time"
)

func Test_func(t *testing.T) {
	var f *os.File
	var err error
	if f, err = os.Open("1.mp4"); err != nil {
		t.Error(err)
		return
	}
	defer f.Close()

	var m *MP4
	if m, err = Decode(f); err != nil {
		t.Error(err)
		return
	}

	ft := Clip((0*60+30)*time.Second, (0*60+40)*time.Second)

	var fw *os.File
	os.Remove("2.mp4")
	if fw, err = os.OpenFile("2.mp4", os.O_CREATE|os.O_WRONLY, 0666); err != nil {
		t.Error(err)
		return
	}

	defer fw.Close()

	if err = EncodeFiltered(fw, m, ft); err != nil {
		t.Error(err)
		return
	}
}
