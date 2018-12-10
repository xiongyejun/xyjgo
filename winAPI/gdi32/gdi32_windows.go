package gdi32

import (
	"syscall"

	"github.com/xiongyejun/xyjgo/winAPI/win"
)

var (
	lib      uintptr
	getPixel uintptr
)

func init() {
	// Library
	lib = win.MustLoadLibrary("gdi32.dll")

	// Functions
	getPixel = win.MustGetProcAddress(lib, "GetPixel")
}

// COLORREF GetPixel(HDC hdc, int nXPos, int nYPos)
func GetPixel(HDC uint32, x, y int) int32 {
	ret, _, _ := syscall.Syscall(getPixel, 3,
		uintptr(HDC),
		uintptr(x),
		uintptr(y))

	return int32(ret)
}

func Free() {
	syscall.FreeLibrary(syscall.Handle(lib))
}
