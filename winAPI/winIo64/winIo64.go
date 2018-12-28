package winIo64

import (
	"syscall"

	"github.com/xiongyejun/xyjgo/winAPI/win"
)

var (
	lib             uintptr
	initializeWinIo uintptr
)

func init() {
	// Library
	lib = win.MustLoadLibrary("WinIo64.dll")

	// Functions
	initializeWinIo = win.MustGetProcAddress(lib, "InitializeWinIo")
}

func InitializeWinIo() uint32 {
	ret, _, _ := syscall.Syscall(initializeWinIo, 0,
		0,
		0,
		0)

	return uint32(ret)
}
