package user32

import (
	"fmt"
	"testing"

	"github.com/xiongyejun/xyjgo/winAPI/win"
)

func Test_func(t *testing.T) {
	defer Free()

	EnumChildWindows(1049218, enumProc)

}

func enumProc(hwnd win.HWND, lParam int) int {
	fmt.Printf("%d  %s\n", hwnd, GetWindowText(hwnd))

	return 1
}
