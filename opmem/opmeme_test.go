package opmem

import (
	"testing"
)

func Test_func(t *testing.T) {
	ret, err := NewByWindow(`C:\Users\Administrator\Documents\08-go\src\github.com\xiongyejun\xyjgo\opmem\exe32\exe32.exe`)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ret.Process)

	retv, err := ret.Scan()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v\n", retv[0].MemInfo)

	b, err := ret.Read(0x1240a088, 4)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("% x\n", b)
}
