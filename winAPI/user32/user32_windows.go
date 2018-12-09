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
	findWindow          uintptr
	getWindowTextLength uintptr
	sendMessage         uintptr
	postMessage         uintptr
	mapVirtualKey       uintptr
	blockInput          uintptr
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
	findWindow = win.MustGetProcAddress(lib, "FindWindowW")
	getWindowTextLength = win.MustGetProcAddress(lib, "GetWindowTextLengthW")
	sendMessage = win.MustGetProcAddress(lib, "SendMessageW")
	postMessage = win.MustGetProcAddress(lib, "PostMessageW")
	mapVirtualKey = win.MustGetProcAddress(lib, "MapVirtualKeyW")
	blockInput = win.MustGetProcAddress(lib, "BlockInput")
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

const (
	WM_KEYDOWN = 0x0100
	WM_KEYUP   = 0x0101
	WM_CHAR    = 0x102
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

//WM_KEYDOWN和WM_KEYUP的 wParam就是虚拟键码，MSDN上可以查到
//也可以通过VkKeyScan将一个字符转换成虚拟键码和shift状态的结合。
//lParam的0到15位为该键在键盘上的重复次数，经常设为1，即按键1次；
//16至23位为键盘的扫描码，通过MapVirtualKey配合其参数可以得到；
//24位为扩展键，即某些右ALT和CTRL；29、30、31位按照说明设置即可
//（第30位对于keydown在和shift等结合的时候通常要设置为1）。
// RESULT SendMessage（HWND hWnd，UINT Msg，WPARAM wParam，LPARAM IParam）
func SendMessage(hWnd uint32, Msg uint, wParam uint16, IParam uint32) uint32 {
	ret, _, _ := syscall.Syscall6(sendMessage, 4,
		uintptr(hWnd),
		uintptr(Msg),
		uintptr(wParam),
		uintptr(IParam),
		0,
		0)

	return uint32(ret)
}
func PostMessage(hWnd uint32, Msg uint, wParam uint16, IParam uint32) uint32 {
	ret, _, _ := syscall.Syscall6(postMessage, 4,
		uintptr(hWnd),
		uintptr(Msg),
		uintptr(wParam),
		uintptr(IParam),
		0,
		0)

	return uint32(ret)
}

// UINT MapVirtualKey（UINT uCode，UINT uMapType）
//uCode：定义一个键的扫描码或虚拟键码。该值如何解释依赖于uMapType参数的值。
//uMapType：定义将要执行的翻译。该参数的值依赖于uCode参数的值。取值如下：
//0：代表uCode是一虚拟键码且被翻译为一扫描码。若一虚拟键码不区分左右，则返回左键的扫描码。若未进行翻译，则函数返回O。
//1：代表uCode是一扫描码且被翻译为一虚拟键码，且此虚拟键码不区分左右。若未进行翻译，则函数返回0。
//2：代表uCode为一虚拟键码且被翻译为一未被移位的字符值存放于返回值的低序字中。死键（发音符号）则通过设置返回值的最高位来表示。若未进行翻译，则函数返回0。
//3：代表uCode为一扫描码且被翻译为区分左右键的一虚拟键码。若未进行翻译，则函数返回0。
//返回值：返回值可以是一扫描码，或一虚拟键码，或一字符值，这完全依赖于不同的uCode和uMapType的值。若未进行翻译，则函数返回0。
func MapVirtualKey(uCode, uMapType uint32) uint32 {
	ret, _, _ := syscall.Syscall(mapVirtualKey, 2,
		uintptr(uCode),
		uintptr(uMapType),
		0)

	return uint32(ret)
}

//BOOL BlockInput( BOOL fBlockIt);
func BlockInput(fBlockIt uint32) uint32 {
	ret, _, _ := syscall.Syscall(blockInput, 1,
		uintptr(fBlockIt),
		0,
		0)

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

// HWND FindWindow（LPCTSTR IpClassName，LPCTSTR IpWindowName）

func FindWindow(IpClassName string, IpWindowName string) uint32 {
	ret, _, _ := syscall.Syscall(findWindow, 2,
		win.StrPtr(IpClassName, win.CODE_UTF16),
		win.StrPtr(IpWindowName, win.CODE_UTF16),
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
