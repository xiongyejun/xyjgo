package syslistview32

import (
	"testing"
	"time"

	"github.com/xiongyejun/xyjgo/winAPI/user32"
)

func Test_getitem(t *testing.T) {
	time.Sleep(2 * time.Second)

	hwnd := user32.GetForegroundWindow()
	t.Logf("%x\r\n", hwnd)

	l, err := NewListView32(hwnd, GetWin7WinStruct())
	if err != nil {
		t.Log(err)
		return
	}

	count := l.GetItemsCount()
	t.Log(count)

	selectindex := l.GetSlectedItemIndex()
	t.Log(selectindex)

	str, err := l.GetItemString(selectindex)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(str)

}
