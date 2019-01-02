package kernel32

import (
	"syscall"
	"unsafe"

	"github.com/xiongyejun/xyjgo/winAPI/win"
)

var (
	lib                uintptr
	getShortPathName   uintptr
	globalAlloc        uintptr
	globalFree         uintptr
	globalLock         uintptr
	globalUnlock       uintptr
	openProcess        uintptr
	virtualAllocEx     uintptr
	virtualFreeEx      uintptr
	writeProcessMemory uintptr
	readProcessMemory  uintptr
	closeHandle        uintptr
	getLastError       uintptr
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
	openProcess = win.MustGetProcAddress(lib, "OpenProcess")
	virtualAllocEx = win.MustGetProcAddress(lib, "VirtualAllocEx")
	virtualFreeEx = win.MustGetProcAddress(lib, "VirtualFreeEx")
	writeProcessMemory = win.MustGetProcAddress(lib, "WriteProcessMemory")
	readProcessMemory = win.MustGetProcAddress(lib, "ReadProcessMemory")
	closeHandle = win.MustGetProcAddress(lib, "CloseHandle")
	getLastError = win.MustGetProcAddress(lib, "GetLastError")
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

const (
	SYNCHRONIZE              = 0x100000
	STANDARD_RIGHTS_REQUIRED = 0xF0000
	PROCESS_ALL_ACCESS       = (STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0xFFF)
	PROCESS_VM_OPERATION     = 0x8
	PROCESS_VM_READ          = 0x10
	PROCESS_VM_WRITE         = 0x20
)

func OpenProcess(dwDesiredAccess, bInheritHandle, dwProcessId uint32) uint32 {
	ret, _, _ := syscall.Syscall(openProcess, 3,
		uintptr(dwDesiredAccess),
		uintptr(bInheritHandle),
		uintptr(dwProcessId))

	return uint32(ret)
}
func CloseHandle(hObject uint32) bool {
	ret, _, _ := syscall.Syscall(closeHandle, 1,
		uintptr(hObject),
		0,
		0)

	return ret != 0
}
func VirtualAllocEx(hProcess uint32, lpAddress uintptr, dwSize uintptr, flAllocationType, flProtect uint32) uintptr {
	ret, _, _ := syscall.Syscall6(virtualAllocEx, 5,
		uintptr(hProcess),
		lpAddress,
		dwSize,
		uintptr(flAllocationType),
		uintptr(flProtect),
		0)

	return uintptr(ret)
}
func VirtualFreeEx(hProcess uint32, lpAddress uintptr, dwSize uintptr, dwFreeType uint32) int32 {
	ret, _, _ := syscall.Syscall6(virtualFreeEx, 4,
		uintptr(hProcess),
		lpAddress,
		dwSize,
		uintptr(dwFreeType),
		0,
		0)

	return int32(ret)
}
func WriteProcessMemory(hProcess uint32, lpBaseAddress uintptr, lpBuffer uintptr, nSize, lpNumberOfBytesWrittenu int32) int32 {
	ret, _, _ := syscall.Syscall6(writeProcessMemory, 5,
		uintptr(hProcess),
		lpBaseAddress,
		lpBuffer,
		uintptr(nSize),
		uintptr(lpNumberOfBytesWrittenu),
		0)

	return int32(ret)
}

func ReadProcessMemory(hProcess uint32, lpBaseAddress uintptr, lpBuffer uintptr, nSize, lpNumberOfBytesWrittenu int32) int32 {
	ret, _, _ := syscall.Syscall6(readProcessMemory, 5,
		uintptr(hProcess),
		lpBaseAddress,
		lpBuffer,
		uintptr(nSize),
		uintptr(lpNumberOfBytesWrittenu),
		0)

	return int32(ret)
}

func GetLastError() uint32 {
	ret, _, _ := syscall.Syscall(getLastError, 0,
		0,
		0,
		0)

	return uint32(ret)
}

func Free() {
	syscall.FreeLibrary(syscall.Handle(lib))
}
