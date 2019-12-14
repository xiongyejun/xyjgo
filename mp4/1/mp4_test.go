package mp4

import (
	"os"
	"testing"
	"time"
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

	//	for i := range m.boxes {
	//		t.Logf("%d type=%s, size=%d\n", i, m.boxes[i].Type(), m.boxes[i].Size())
	//	}
	//	t.Logf("%d\n", len(m.Moov.Trak))
	//	return

	ft := Clip((0*60+5)*time.Second, 10*time.Second)
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
