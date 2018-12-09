package keyboard

import (
	"testing"

	"github.com/xiongyejun/xyjgo/winAPI/user32"
)

func Test_press(t *testing.T) {
	//	hwnd := user32.FindWindow("", "1.txt - 记事本")
	hwnd := user32.FindWindow("", "MapleStory")
	t.Log("hwnd=", hwnd)

	//	ret := SendMessage(hwnd, VK_O)
	ret := PostMessage(hwnd, VK_P)

	t.Log(ret)
}
