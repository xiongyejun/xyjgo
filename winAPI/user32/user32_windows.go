package user32

import (
	"syscall"
	"unsafe"

	"github.com/xiongyejun/xyjgo/winAPI/win"
)

var (
	lib                 uintptr
	sendInput           uintptr
	keybd_event         uintptr
	getForegroundWindow uintptr
	getWindowText       uintptr
	getWindowTextLength uintptr
)

//Private Declare Function mciSendStringA Lib "winmm.dll" _
//        (ByVal lpstrCommand As String, ByVal lpstrReturnString As String,
//        ByVal uReturnLength As Integer, ByVal hwndCallback As Integer) As Integer

func init() {
	// Library
	lib = win.MustLoadLibrary("user32.dll")

	// Functions
	sendInput = win.MustGetProcAddress(lib, "SendInput")
	keybd_event = win.MustGetProcAddress(lib, "keybd_event")
	getForegroundWindow = win.MustGetProcAddress(lib, "GetForegroundWindow")
	getWindowText = win.MustGetProcAddress(lib, "GetWindowTextW")
	getWindowTextLength = win.MustGetProcAddress(lib, "GetWindowTextLengthW")
}

// INPUT Type
const (
	INPUT_MOUSE    = 0
	INPUT_KEYBOARD = 1
	INPUT_HARDWARE = 2
)

type MOUSE_INPUT struct {
	Type uint32
	Mi   MOUSEINPUT
}

// MOUSEINPUT MouseData
const (
	XBUTTON1 = 0x0001
	XBUTTON2 = 0x0002
)

// MOUSEINPUT DwFlags
const (
	MOUSEEVENTF_ABSOLUTE        = 0x8000
	MOUSEEVENTF_HWHEEL          = 0x1000
	MOUSEEVENTF_MOVE            = 0x0001
	MOUSEEVENTF_MOVE_NOCOALESCE = 0x2000
	MOUSEEVENTF_LEFTDOWN        = 0x0002
	MOUSEEVENTF_LEFTUP          = 0x0004
	MOUSEEVENTF_RIGHTDOWN       = 0x0008
	MOUSEEVENTF_RIGHTUP         = 0x0010
	MOUSEEVENTF_MIDDLEDOWN      = 0x0020
	MOUSEEVENTF_MIDDLEUP        = 0x0040
	MOUSEEVENTF_VIRTUALDESK     = 0x4000
	MOUSEEVENTF_WHEEL           = 0x0800
	MOUSEEVENTF_XDOWN           = 0x0080
	MOUSEEVENTF_XUP             = 0x0100
)

// KEYBDINPUT DwFlags
const (
	KEYEVENTF_EXTENDEDKEY = 0x0001
	KEYEVENTF_KEYUP       = 0x0002
	KEYEVENTF_SCANCODE    = 0x0008
	KEYEVENTF_UNICODE     = 0x0004
)

type MOUSEINPUT struct {
	Dx          int32
	Dy          int32
	MouseData   uint32
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
}

type KEYBD_INPUT struct {
	Type uint32
	Ki   KEYBDINPUT
}

type KEYBDINPUT struct {
	WVk         uint16
	WScan       uint16
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
	Unused      [8]byte
}

type HARDWARE_INPUT struct {
	Type uint32
	Hi   HARDWAREINPUT
}

type HARDWAREINPUT struct {
	UMsg    uint32
	WParamL uint16
	WParamH uint16
	Unused  [16]byte
}

//UINT Send Input{
//	UINT nInput;
//	LPINPUT pInput;
//	INT cbSize;
//}

// pInputs expects a unsafe.Pointer to a slice of MOUSE_INPUT or KEYBD_INPUT or HARDWARE_INPUT structs.
func SendInput(nInputs uint32, pInputs unsafe.Pointer, cbSize int32) uint32 {
	ret, _, _ := syscall.Syscall(sendInput, 3,
		uintptr(nInputs),
		uintptr(pInputs),
		uintptr(cbSize))

	return uint32(ret)
}

//byte bVk,    //虚拟键值
//byte bScan,// 一般为0
//int dwFlags,  //这里是整数类型  0 为按下，2为释放
//int dwExtraInfo  //这里是整数类型 一般情况下设成为 0
func Keybd_event(bVk byte, bScan byte, dwFlags int, dwExtraInfo int) uint32 {
	ret, _, _ := syscall.Syscall6(keybd_event, 4,
		uintptr(bVk),
		uintptr(bScan),
		uintptr(dwFlags),
		uintptr(dwExtraInfo),
		0,
		0)

	return uint32(ret)
}

func GetForegroundWindow() uint32 {
	ret, _, _ := syscall.Syscall(getForegroundWindow, 0,
		0,
		0,
		0)

	return uint32(ret)
}

func GetWindowText(hwnd uint32) string {
	iLen := GetWindowTextLength(hwnd) + 1
	buf := make([]uint16, iLen)

	_, _, err := syscall.Syscall(getWindowText, 3,
		uintptr(hwnd), uintptr(unsafe.Pointer(&buf[0])),
		uintptr(iLen))

	if err > 0 {
		return ""
	}
	return syscall.UTF16ToString(buf)

}

func GetWindowTextLength(hwnd uint32) uint32 {
	ret, _, _ := syscall.Syscall(getWindowTextLength, 1, uintptr(hwnd), 0, 0)
	return uint32(ret)
}

func Free() {
	syscall.FreeLibrary(syscall.Handle(lib))
}
