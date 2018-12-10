package user32

import (
	"testing"
)

func Test_func(t *testing.T) {
	defer Free()

	hwnd := FindWindow("", "1.txt - 记事本")
	t.Log(hwnd)

	hdc := GetDC(hwnd)
	t.Log(hdc)

	ret := ReleaseDC(hwnd, hdc)
	t.Log(ret)
}
