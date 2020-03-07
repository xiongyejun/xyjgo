package user32

import (
	"syscall"
	"unsafe"

	"github.com/xiongyejun/xyjgo/winAPI/win"
)

var (
	lib                      uintptr
	sendInput                uintptr
	keybd_event              uintptr
	getForegroundWindow      uintptr
	getWindowText            uintptr
	findWindow               uintptr
	getWindowTextLength      uintptr
	sendMessage              uintptr
	postMessage              uintptr
	mapVirtualKey            uintptr
	blockInput               uintptr
	getDC                    uintptr
	releaseDC                uintptr
	getWindowRect            uintptr
	getClientRect            uintptr
	getWindowThreadProcessId uintptr
	findWindowEx             uintptr
	messageBox               uintptr
	getActiveWindow          uintptr
	getWindowLong            uintptr
	setWindowLong            uintptr
	getWindowLongPtr         uintptr
	setWindowLongPtr         uintptr
	updateWindow             uintptr
	showWindow               uintptr
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
	findWindowEx = win.MustGetProcAddress(lib, "FindWindowExW")
	getWindowTextLength = win.MustGetProcAddress(lib, "GetWindowTextLengthW")
	sendMessage = win.MustGetProcAddress(lib, "SendMessageW")
	postMessage = win.MustGetProcAddress(lib, "PostMessageW")
	mapVirtualKey = win.MustGetProcAddress(lib, "MapVirtualKeyW")
	blockInput = win.MustGetProcAddress(lib, "BlockInput")
	getDC = win.MustGetProcAddress(lib, "GetDC")
	releaseDC = win.MustGetProcAddress(lib, "ReleaseDC")
	getWindowRect = win.MustGetProcAddress(lib, "GetWindowRect")
	getClientRect = win.MustGetProcAddress(lib, "GetClientRect")
	getWindowThreadProcessId = win.MustGetProcAddress(lib, "GetWindowThreadProcessId")
	messageBox = win.MustGetProcAddress(lib, "MessageBoxW")
	getActiveWindow = win.MustGetProcAddress(lib, "GetActiveWindow")
	getWindowLong = win.MustGetProcAddress(lib, "GetWindowLongW")
	setWindowLong = win.MustGetProcAddress(lib, "SetWindowLongW")
	//	getWindowLongPtr = win.MustGetProcAddress(lib, "GetWindowLongPtrW")
	//	setWindowLongPtr = win.MustGetProcAddress(lib, "SetWindowLongPtrW")
	updateWindow = win.MustGetProcAddress(lib, "UpdateWindow")
	showWindow = win.MustGetProcAddress(lib, "ShowWindow")
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

// Window message constants
const (
	WM_APP                    = 32768
	WM_ACTIVATE               = 6
	WM_ACTIVATEAPP            = 28
	WM_AFXFIRST               = 864
	WM_AFXLAST                = 895
	WM_ASKCBFORMATNAME        = 780
	WM_CANCELJOURNAL          = 75
	WM_CANCELMODE             = 31
	WM_CAPTURECHANGED         = 533
	WM_CHANGECBCHAIN          = 781
	WM_CHAR                   = 258
	WM_CHARTOITEM             = 47
	WM_CHILDACTIVATE          = 34
	WM_CLEAR                  = 771
	WM_CLOSE                  = 16
	WM_COMMAND                = 273
	WM_COMMNOTIFY             = 68 /* OBSOLETE */
	WM_COMPACTING             = 65
	WM_COMPAREITEM            = 57
	WM_CONTEXTMENU            = 123
	WM_COPY                   = 769
	WM_COPYDATA               = 74
	WM_CREATE                 = 1
	WM_CTLCOLORBTN            = 309
	WM_CTLCOLORDLG            = 310
	WM_CTLCOLOREDIT           = 307
	WM_CTLCOLORLISTBOX        = 308
	WM_CTLCOLORMSGBOX         = 306
	WM_CTLCOLORSCROLLBAR      = 311
	WM_CTLCOLORSTATIC         = 312
	WM_CUT                    = 768
	WM_DEADCHAR               = 259
	WM_DELETEITEM             = 45
	WM_DESTROY                = 2
	WM_DESTROYCLIPBOARD       = 775
	WM_DEVICECHANGE           = 537
	WM_DEVMODECHANGE          = 27
	WM_DISPLAYCHANGE          = 126
	WM_DRAWCLIPBOARD          = 776
	WM_DRAWITEM               = 43
	WM_DROPFILES              = 563
	WM_ENABLE                 = 10
	WM_ENDSESSION             = 22
	WM_ENTERIDLE              = 289
	WM_ENTERMENULOOP          = 529
	WM_ENTERSIZEMOVE          = 561
	WM_ERASEBKGND             = 20
	WM_EXITMENULOOP           = 530
	WM_EXITSIZEMOVE           = 562
	WM_FONTCHANGE             = 29
	WM_GETDLGCODE             = 135
	WM_GETFONT                = 49
	WM_GETHOTKEY              = 51
	WM_GETICON                = 127
	WM_GETMINMAXINFO          = 36
	WM_GETTEXT                = 13
	WM_GETTEXTLENGTH          = 14
	WM_HANDHELDFIRST          = 856
	WM_HANDHELDLAST           = 863
	WM_HELP                   = 83
	WM_HOTKEY                 = 786
	WM_HSCROLL                = 276
	WM_HSCROLLCLIPBOARD       = 782
	WM_ICONERASEBKGND         = 39
	WM_INITDIALOG             = 272
	WM_INITMENU               = 278
	WM_INITMENUPOPUP          = 279
	WM_INPUT                  = 0X00FF
	WM_INPUTLANGCHANGE        = 81
	WM_INPUTLANGCHANGEREQUEST = 80
	WM_KEYDOWN                = 256
	WM_KEYUP                  = 257
	WM_KILLFOCUS              = 8
	WM_MDIACTIVATE            = 546
	WM_MDICASCADE             = 551
	WM_MDICREATE              = 544
	WM_MDIDESTROY             = 545
	WM_MDIGETACTIVE           = 553
	WM_MDIICONARRANGE         = 552
	WM_MDIMAXIMIZE            = 549
	WM_MDINEXT                = 548
	WM_MDIREFRESHMENU         = 564
	WM_MDIRESTORE             = 547
	WM_MDISETMENU             = 560
	WM_MDITILE                = 550
	WM_MEASUREITEM            = 44
	WM_GETOBJECT              = 0X003D
	WM_CHANGEUISTATE          = 0X0127
	WM_UPDATEUISTATE          = 0X0128
	WM_QUERYUISTATE           = 0X0129
	WM_UNINITMENUPOPUP        = 0X0125
	WM_MENURBUTTONUP          = 290
	WM_MENUCOMMAND            = 0X0126
	WM_MENUGETOBJECT          = 0X0124
	WM_MENUDRAG               = 0X0123
	WM_APPCOMMAND             = 0X0319
	WM_MENUCHAR               = 288
	WM_MENUSELECT             = 287
	WM_MOVE                   = 3
	WM_MOVING                 = 534
	WM_NCACTIVATE             = 134
	WM_NCCALCSIZE             = 131
	WM_NCCREATE               = 129
	WM_NCDESTROY              = 130
	WM_NCHITTEST              = 132
	WM_NCLBUTTONDBLCLK        = 163
	WM_NCLBUTTONDOWN          = 161
	WM_NCLBUTTONUP            = 162
	WM_NCMBUTTONDBLCLK        = 169
	WM_NCMBUTTONDOWN          = 167
	WM_NCMBUTTONUP            = 168
	WM_NCXBUTTONDOWN          = 171
	WM_NCXBUTTONUP            = 172
	WM_NCXBUTTONDBLCLK        = 173
	WM_NCMOUSEHOVER           = 0X02A0
	WM_NCMOUSELEAVE           = 0X02A2
	WM_NCMOUSEMOVE            = 160
	WM_NCPAINT                = 133
	WM_NCRBUTTONDBLCLK        = 166
	WM_NCRBUTTONDOWN          = 164
	WM_NCRBUTTONUP            = 165
	WM_NEXTDLGCTL             = 40
	WM_NEXTMENU               = 531
	WM_NOTIFY                 = 78
	WM_NOTIFYFORMAT           = 85
	WM_NULL                   = 0
	WM_PAINT                  = 15
	WM_PAINTCLIPBOARD         = 777
	WM_PAINTICON              = 38
	WM_PALETTECHANGED         = 785
	WM_PALETTEISCHANGING      = 784
	WM_PARENTNOTIFY           = 528
	WM_PASTE                  = 770
	WM_PENWINFIRST            = 896
	WM_PENWINLAST             = 911
	WM_POWER                  = 72
	WM_POWERBROADCAST         = 536
	WM_PRINT                  = 791
	WM_PRINTCLIENT            = 792
	WM_QUERYDRAGICON          = 55
	WM_QUERYENDSESSION        = 17
	WM_QUERYNEWPALETTE        = 783
	WM_QUERYOPEN              = 19
	WM_QUEUESYNC              = 35
	WM_QUIT                   = 18
	WM_RENDERALLFORMATS       = 774
	WM_RENDERFORMAT           = 773
	WM_SETCURSOR              = 32
	WM_SETFOCUS               = 7
	WM_SETFONT                = 48
	WM_SETHOTKEY              = 50
	WM_SETICON                = 128
	WM_SETREDRAW              = 11
	WM_SETTEXT                = 12
	WM_SETTINGCHANGE          = 26
	WM_SHOWWINDOW             = 24
	WM_SIZE                   = 5
	WM_SIZECLIPBOARD          = 779
	WM_SIZING                 = 532
	WM_SPOOLERSTATUS          = 42
	WM_STYLECHANGED           = 125
	WM_STYLECHANGING          = 124
	WM_SYSCHAR                = 262
	WM_SYSCOLORCHANGE         = 21
	WM_SYSCOMMAND             = 274
	WM_SYSDEADCHAR            = 263
	WM_SYSKEYDOWN             = 260
	WM_SYSKEYUP               = 261
	WM_TCARD                  = 82
	WM_THEMECHANGED           = 794
	WM_TIMECHANGE             = 30
	WM_TIMER                  = 275
	WM_UNDO                   = 772
	WM_USER                   = 1024
	WM_USERCHANGED            = 84
	WM_VKEYTOITEM             = 46
	WM_VSCROLL                = 277
	WM_VSCROLLCLIPBOARD       = 778
	WM_WINDOWPOSCHANGED       = 71
	WM_WINDOWPOSCHANGING      = 70
	WM_WININICHANGE           = 26
	WM_KEYFIRST               = 256
	WM_KEYLAST                = 264
	WM_SYNCPAINT              = 136
	WM_MOUSEACTIVATE          = 33
	WM_MOUSEMOVE              = 512
	WM_LBUTTONDOWN            = 513
	WM_LBUTTONUP              = 514
	WM_LBUTTONDBLCLK          = 515
	WM_RBUTTONDOWN            = 516
	WM_RBUTTONUP              = 517
	WM_RBUTTONDBLCLK          = 518
	WM_MBUTTONDOWN            = 519
	WM_MBUTTONUP              = 520
	WM_MBUTTONDBLCLK          = 521
	WM_MOUSEWHEEL             = 522
	WM_MOUSEFIRST             = 512
	WM_XBUTTONDOWN            = 523
	WM_XBUTTONUP              = 524
	WM_XBUTTONDBLCLK          = 525
	WM_MOUSELAST              = 525
	WM_MOUSEHOVER             = 0X2A1
	WM_MOUSELEAVE             = 0X2A3
	WM_CLIPBOARDUPDATE        = 0x031D
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

type RECT struct {
	Left, Top, Right, Bottom int32
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
func SendMessage(hWnd uint32, Msg uint, wParam uintptr, IParam uintptr) uint32 {
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

func GetActiveWindow() uint32 {
	ret, _, _ := syscall.Syscall(getActiveWindow, 0,
		0,
		0,
		0)

	return uint32(ret)
}

// HDC GetDC(HWND hWnd)；
func GetDC(HWND uint32) uint32 {
	ret, _, _ := syscall.Syscall(getDC, 1,
		uintptr(HWND),
		0,
		0)

	return uint32(ret)
}

// int ReleaseDC(HWND hWnd, HDC hdc)；
func ReleaseDC(HWND uint32, HDC uint32) uint32 {
	ret, _, _ := syscall.Syscall(releaseDC, 2,
		uintptr(HWND),
		uintptr(HDC),
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

func FindWindowEx(hwndParent, hwndChildAfter uint32, IpClassName string, IpWindowName string) uint32 {
	ret, _, _ := syscall.Syscall6(findWindowEx, 4,
		uintptr(hwndParent),
		uintptr(hwndChildAfter),
		win.StrPtr(IpClassName, win.CODE_UTF16),
		win.StrPtr(IpWindowName, win.CODE_UTF16),
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

func GetWindowRect(hWnd uint32, rect *RECT) bool {
	ret, _, _ := syscall.Syscall(getWindowRect, 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(rect)),
		0)

	return ret != 0
}
func GetClientRect(hWnd uint32, rect *RECT) bool {
	ret, _, _ := syscall.Syscall(getClientRect, 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(rect)),
		0)

	return ret != 0
}

func GetWindowThreadProcessId(hwnd uint32, pid *uint32) uint32 {
	ret, _, _ := syscall.Syscall(getWindowThreadProcessId, 2,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(pid)),
		0)

	return uint32(ret)
}

const (
	MB_OK                = 0x00000000
	MB_OKCANCEL          = 0x00000001
	MB_ABORTRETRYIGNORE  = 0x00000002
	MB_YESNOCANCEL       = 0x00000003
	MB_YESNO             = 0x00000004
	MB_RETRYCANCEL       = 0x00000005
	MB_CANCELTRYCONTINUE = 0x00000006
	MB_ICONHAND          = 0x00000010
	MB_ICONQUESTION      = 0x00000020
	MB_ICONEXCLAMATION   = 0x00000030
	MB_ICONASTERISK      = 0x00000040
	MB_USERICON          = 0x00000080
	MB_ICONWARNING       = MB_ICONEXCLAMATION
	MB_ICONERROR         = MB_ICONHAND
	MB_ICONINFORMATION   = MB_ICONASTERISK
	MB_ICONSTOP          = MB_ICONHAND
	MB_DEFBUTTON1        = 0x00000000
	MB_DEFBUTTON2        = 0x00000100
	MB_DEFBUTTON3        = 0x00000200
	MB_DEFBUTTON4        = 0x00000300

	MB_RETURN_YES = 0x00000006
	MB_RETURN_NO  = 0x00000007
)

func MessageBox(hWnd uint32, lpText, lpCaption string, uType uint32) int32 {
	ret, _, _ := syscall.Syscall6(messageBox, 4,
		uintptr(hWnd),
		win.StrPtr(lpText, win.CODE_UTF16),
		win.StrPtr(lpCaption, win.CODE_UTF16),
		uintptr(uType),
		0,
		0)

	return int32(ret)
}

// GetWindowLong and GetWindowLongPtr constants
const (
	GWL_EXSTYLE     = -20
	GWL_STYLE       = -16
	GWL_WNDPROC     = -4
	GWLP_WNDPROC    = -4
	GWL_HINSTANCE   = -6
	GWLP_HINSTANCE  = -6
	GWL_HWNDPARENT  = -8
	GWLP_HWNDPARENT = -8
	GWL_ID          = -12
	GWLP_ID         = -12
	GWL_USERDATA    = -21
	GWLP_USERDATA   = -21
)

func GetWindowLong(hWnd uint32, index int32) int32 {
	ret, _, _ := syscall.Syscall(getWindowLong, 2,
		uintptr(hWnd),
		uintptr(index),
		0)

	return int32(ret)
}

// Extended window style constants
const (
	WS_EX_DLGMODALFRAME    = 0X00000001
	WS_EX_NOPARENTNOTIFY   = 0X00000004
	WS_EX_TOPMOST          = 0X00000008
	WS_EX_ACCEPTFILES      = 0X00000010
	WS_EX_TRANSPARENT      = 0X00000020
	WS_EX_MDICHILD         = 0X00000040
	WS_EX_TOOLWINDOW       = 0X00000080
	WS_EX_WINDOWEDGE       = 0X00000100
	WS_EX_CLIENTEDGE       = 0X00000200
	WS_EX_CONTEXTHELP      = 0X00000400
	WS_EX_RIGHT            = 0X00001000
	WS_EX_LEFT             = 0X00000000
	WS_EX_RTLREADING       = 0X00002000
	WS_EX_LTRREADING       = 0X00000000
	WS_EX_LEFTSCROLLBAR    = 0X00004000
	WS_EX_RIGHTSCROLLBAR   = 0X00000000
	WS_EX_CONTROLPARENT    = 0X00010000
	WS_EX_STATICEDGE       = 0X00020000
	WS_EX_APPWINDOW        = 0X00040000
	WS_EX_OVERLAPPEDWINDOW = 0X00000100 | 0X00000200
	WS_EX_PALETTEWINDOW    = 0X00000100 | 0X00000080 | 0X00000008
	WS_EX_LAYERED          = 0X00080000
	WS_EX_NOINHERITLAYOUT  = 0X00100000
	WS_EX_LAYOUTRTL        = 0X00400000
	WS_EX_COMPOSITED       = 0X02000000
	WS_EX_NOACTIVATE       = 0X08000000
)

func SetWindowLong(hWnd uint32, index, value int32) int32 {
	ret, _, _ := syscall.Syscall(setWindowLong, 3,
		uintptr(hWnd),
		uintptr(index),
		uintptr(value))

	return int32(ret)
}

func SetWindowLongPtr(hWnd uint32, index int, value uintptr) uintptr {
	ret, _, _ := syscall.Syscall(setWindowLongPtr, 3,
		uintptr(hWnd),
		uintptr(index),
		value)

	return ret
}
func GetWindowLongPtr(hWnd uint32, index int32) uintptr {
	ret, _, _ := syscall.Syscall(getWindowLongPtr, 2,
		uintptr(hWnd),
		uintptr(index),
		0)

	return ret
}
func UpdateWindow(hwnd uint32) bool {
	ret, _, _ := syscall.Syscall(updateWindow, 1,
		uintptr(hwnd),
		0,
		0)

	return ret != 0
}

// ShowWindow constants
const (
	SW_HIDE            = 0
	SW_NORMAL          = 1
	SW_SHOWNORMAL      = 1
	SW_SHOWMINIMIZED   = 2
	SW_MAXIMIZE        = 3
	SW_SHOWMAXIMIZED   = 3
	SW_SHOWNOACTIVATE  = 4
	SW_SHOW            = 5
	SW_MINIMIZE        = 6
	SW_SHOWMINNOACTIVE = 7
	SW_SHOWNA          = 8
	SW_RESTORE         = 9
	SW_SHOWDEFAULT     = 10
	SW_FORCEMINIMIZE   = 11
)

func ShowWindow(hWnd uint32, nCmdShow int32) bool {
	ret, _, _ := syscall.Syscall(showWindow, 2,
		uintptr(hWnd),
		uintptr(nCmdShow),
		0)

	return ret != 0
}
func Free() {
	syscall.FreeLibrary(syscall.Handle(lib))
}
