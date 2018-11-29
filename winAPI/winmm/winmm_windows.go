package winmm

import (
	"bytes"
	"syscall"
	"unsafe"

	"github.com/axgle/mahonia"

	"github.com/xiongyejun/xyjgo/winAPI/win"
)

var (
	libwinmm          uintptr
	mciSendStringA    uintptr
	mciGetErrorString uintptr
)

//Private Declare Function mciSendStringA Lib "winmm.dll" _
//        (ByVal lpstrCommand As String, ByVal lpstrReturnString As String,
//        ByVal uReturnLength As Integer, ByVal hwndCallback As Integer) As Integer

func init() {
	// Library
	libwinmm = win.MustLoadLibrary("winmm.dll")

	// Functions
	mciSendStringA = win.MustGetProcAddress(libwinmm, "mciSendStringA")
	mciGetErrorString = win.MustGetProcAddress(libwinmm, "mciGetErrorStringA")
}

func MciSendStringA(lpstrCommand, lpstrReturnString string, uReturnLength, hwndCallback uint32) int32 {
	ret, _, _ := syscall.Syscall6(mciSendStringA, 4,
		win.StrPtr(lpstrCommand, win.CODE_GBK),
		win.StrPtr(lpstrReturnString, win.CODE_GBK),
		uintptr(uReturnLength),
		uintptr(hwndCallback),
		0,
		0)

	return int32(ret)
}

// 根据错误的ID，返回错误信息
func GetErrorString(errID int32) string {
	var cchErrorText uint32 = 512
	var lpszErrorText = make([]byte, cchErrorText)
	MciGetErrorString(errID, uintptr(unsafe.Pointer(&lpszErrorText[0])), cchErrorText)
	index := bytes.Index(lpszErrorText, []byte{0})

	decoder := mahonia.NewDecoder("gbk")
	return decoder.ConvertString(string(lpszErrorText[:index]))
}

func MciGetErrorString(fdwError int32, lpszErrorText uintptr, cchErrorText uint32) int32 {
	ret, _, _ := syscall.Syscall6(mciGetErrorString, 3,
		uintptr(fdwError),
		lpszErrorText,
		uintptr(cchErrorText),
		0,
		0,
		0)

	return int32(ret)
}

func Free() {
	syscall.FreeLibrary(syscall.Handle(libwinmm))
}
