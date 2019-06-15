package pic

import (
	"testing"
)

func Test_new(t *testing.T) {
	if p, err := New("03-01.bmp"); err != nil {
		t.Error(err)
		return
	} else {
		t.Log(p.path)
	}

}
