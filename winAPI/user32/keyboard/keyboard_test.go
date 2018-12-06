package keyboard

import (
	"testing"

	"github.com/xiongyejun/xyjgo/winAPI/user32"
)

func Test_press(t *testing.T) {
	//	Press('2')
	//	Press('3')
	//	Press('4')

	hwnd := user32.FindWindow("", "管理员: 命令提示符")
	t.Log(hwnd)

	SendMessage(hwnd, 'A')
	SendMessage(hwnd, '1')
}
