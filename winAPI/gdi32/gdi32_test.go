package gdi32

import (
	"testing"

	"github.com/xiongyejun/xyjgo/winAPI/user32"
)

func Test_func(t *testing.T) {
	defer Free()
	defer user32.Free()

	hwnd := user32.FindWindow("", "1.txt - 记事本")
	t.Log(hwnd)

	hdc := user32.GetDC(hwnd)
	t.Log(hdc)

	pix := GetPixel(hdc, 100, 100)
	t.Log(pix)

	ret := user32.ReleaseDC(hwnd, hdc)
	t.Log(ret)
}
