package advapi32

import (
	"syscall"
	"unsafe"

	"github.com/xiongyejun/xyjgo/winAPI/win"
)

var (
	lib             uintptr
	regOpenKeyEx    uintptr
	regCloseKey     uintptr
	regQueryValueEx uintptr
	regEnumValue    uintptr
	regSetValueEx   uintptr
)

func init() {
	// Library
	lib = win.MustLoadLibrary("advapi32.dll")

	// Functions
	regOpenKeyEx = win.MustGetProcAddress(lib, "RegOpenKeyExW")
	regCloseKey = win.MustGetProcAddress(lib, "RegCloseKey")
	regQueryValueEx = win.MustGetProcAddress(lib, "RegQueryValueExW")
	regEnumValue = win.MustGetProcAddress(lib, "RegEnumValueW")
	regSetValueEx = win.MustGetProcAddress(lib, "RegSetValueExW")
}

const KEY_READ uint32 = 0x20019
const KEY_WRITE uint32 = 0x20006

type HKEY uintptr

const (
	HKEY_CLASSES_ROOT     HKEY = 0x80000000
	HKEY_CURRENT_USER     HKEY = 0x80000001
	HKEY_LOCAL_MACHINE    HKEY = 0x80000002
	HKEY_USERS            HKEY = 0x80000003
	HKEY_PERFORMANCE_DATA HKEY = 0x80000004
	HKEY_CURRENT_CONFIG   HKEY = 0x80000005
	HKEY_DYN_DATA         HKEY = 0x80000006
)

const (
	REG_NONE      uint64 = 0 // No value type
	REG_SZ               = 1 // Unicode nul terminated string
	REG_EXPAND_SZ        = 2 // Unicode nul terminated string
	// (with environment variable references)
	REG_BINARY                     = 3 // Free form binary
	REG_DWORD                      = 4 // 32-bit number
	REG_DWORD_LITTLE_ENDIAN        = 4 // 32-bit number (same as REG_DWORD)
	REG_DWORD_BIG_ENDIAN           = 5 // 32-bit number
	REG_LINK                       = 6 // Symbolic Link (unicode)
	REG_MULTI_SZ                   = 7 // Multiple Unicode strings
	REG_RESOURCE_LIST              = 8 // Resource list in the resource map
	REG_FULL_RESOURCE_DESCRIPTOR   = 9 // Resource list in the hardware description
	REG_RESOURCE_REQUIREMENTS_LIST = 10
	REG_QWORD                      = 11 // 64-bit number
	REG_QWORD_LITTLE_ENDIAN        = 11 // 64-bit number (same as REG_QWORD)

)

func RegOpenKeyEx(hKey HKEY, lpSubKey string, ulOptions uint32, samDesired uint32, phkResult *uint32) int32 {
	ret, _, _ := syscall.Syscall6(regOpenKeyEx, 5,
		uintptr(hKey),
		win.StrPtr(lpSubKey, win.CODE_UTF16),
		uintptr(ulOptions),
		uintptr(samDesired),
		uintptr(unsafe.Pointer(phkResult)),
		0)

	return int32(ret)
}
func RegCloseKey(hKey uint32) int32 {
	ret, _, _ := syscall.Syscall(regCloseKey, 1,
		uintptr(hKey),
		0,
		0)

	return int32(ret)
}

func RegQueryValueEx(hKey HKEY, lpValueName string, lpReserved, lpType *uint32, lpData *byte, lpcbData *uint32) int32 {
	ret, _, _ := syscall.Syscall6(regQueryValueEx, 6,
		uintptr(hKey),
		win.StrPtr(lpValueName, win.CODE_UTF16),
		uintptr(unsafe.Pointer(lpReserved)),
		uintptr(unsafe.Pointer(lpType)),
		uintptr(unsafe.Pointer(lpData)),
		uintptr(unsafe.Pointer(lpcbData)))

	return int32(ret)
}

func RegEnumValue(hKey HKEY, index uint32, lpValueName string, lpcchValueName *uint32, lpReserved, lpType *uint32, lpData *byte, lpcbData *uint32) int32 {
	ret, _, _ := syscall.Syscall9(regEnumValue, 8,
		uintptr(hKey),
		uintptr(index),
		win.StrPtr(lpValueName, win.CODE_UTF16),
		uintptr(unsafe.Pointer(lpcchValueName)),
		uintptr(unsafe.Pointer(lpReserved)),
		uintptr(unsafe.Pointer(lpType)),
		uintptr(unsafe.Pointer(lpData)),
		uintptr(unsafe.Pointer(lpcbData)),
		0)
	return int32(ret)
}

func RegSetValueEx(hKey HKEY, lpValueName string, lpReserved, lpDataType uint64, lpData *byte, cbData uint32) int32 {
	ret, _, _ := syscall.Syscall6(regSetValueEx, 6,
		uintptr(hKey),
		win.StrPtr(lpValueName, win.CODE_UTF16),
		uintptr(lpReserved),
		uintptr(lpDataType),
		uintptr(unsafe.Pointer(lpData)),
		uintptr(cbData))
	return int32(ret)
}

func Free() {
	syscall.FreeLibrary(syscall.Handle(lib))
}
