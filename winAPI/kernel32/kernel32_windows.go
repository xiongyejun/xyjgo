package kernel32

import (
	"syscall"
	"unsafe"

	"github.com/xiongyejun/xyjgo/winAPI/win"
)

var (
	lib              uintptr
	getShortPathName uintptr
	globalAlloc      uintptr
	globalFree       uintptr
	globalLock       uintptr
	globalUnlock     uintptr
)

// GlobalAlloc flags
const (
	GHND          = 0x0042
	GMEM_FIXED    = 0x0000
	GMEM_MOVEABLE = 0x0002
	GMEM_ZEROINIT = 0x0040
	GPTR          = 0x004
)

func init() {
	// Library
	lib = win.MustLoadLibrary("kernel32.dll")

	// Functions
	getShortPathName = win.MustGetProcAddress(lib, "GetShortPathNameA")
	globalAlloc = win.MustGetProcAddress(lib, "GlobalAlloc")
	globalFree = win.MustGetProcAddress(lib, "GlobalFree")
	globalLock = win.MustGetProcAddress(lib, "GlobalLock")
	globalUnlock = win.MustGetProcAddress(lib, "GlobalUnlock")
}

func GlobalAlloc(uFlags uint32, dwBytes uintptr) uint32 {
	ret, _, _ := syscall.Syscall(globalAlloc, 2,
		uintptr(uFlags),
		dwBytes,
		0)

	return uint32(ret)
}

func GlobalFree(hMem uint32) uint32 {
	ret, _, _ := syscall.Syscall(globalFree, 1,
		uintptr(hMem),
		0,
		0)

	return uint32(ret)
}

func GlobalLock(hMem uint32) unsafe.Pointer {
	ret, _, _ := syscall.Syscall(globalLock, 1,
		uintptr(hMem),
		0,
		0)

	return unsafe.Pointer(ret)
}

func GlobalUnlock(hMem uint32) bool {
	ret, _, _ := syscall.Syscall(globalUnlock, 1,
		uintptr(hMem),
		0,
		0)

	return ret != 0
}

func GetShortPathName(longPath string, shortPath uintptr, shortBufferSize uint32) int32 {
	ret, _, _ := syscall.Syscall6(getShortPathName, 3,
		win.StrPtr(longPath, win.CODE_UTF16),
		shortPath,
		uintptr(shortBufferSize),
		0,
		0,
		0)

	return int32(ret)
}

func Free() {
	syscall.FreeLibrary(syscall.Handle(lib))
}
