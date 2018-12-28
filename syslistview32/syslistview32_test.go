package syslistview32

import (
	"testing"
)

func Test_getitem(t *testing.T) {
	l, err := NewListView32(0x00740340)
	if err != nil {
		t.Log(err)
		return
	}
	selectindex := l.GetSlectedItemIndex()

	str, err := l.GetItemString(selectindex)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(str)

	path, err := l.GetFolderPath()
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(path)
}
