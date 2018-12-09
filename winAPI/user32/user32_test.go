package user32

import (
	"testing"
)

func Test_func(t *testing.T) {
	hwnd := FindWindow("", "MapleStory")
	t.Log(hwnd)

	ret := SendMessage(hwnd, WM_KEYDOWN, 'O', MapVirtualKey('O', 0))

	t.Log(ret)
}
