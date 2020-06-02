package user32

import (
	"syscall"
	"unsafe"

	"github.com/xiongyejun/xyjgo/winAPI/win"
)

var (
	lib                      uintptr
	destroyWindow            uintptr
	sendInput                uintptr
	keybd_event              uintptr
	getForegroundWindow      uintptr
	getWindowText            uintptr
	findWindow               uintptr
	getWindowTextLength      uintptr
	sendMessage              uintptr
	postMessage              uintptr
	getMessage               uintptr
	mapVirtualKey            uintptr
	blockInput               uintptr
	getDC                    uintptr
	releaseDC                uintptr
	getWindowRect            uintptr
	getClientRect            uintptr
	getWindowThreadProcessId uintptr
	findWindowEx             uintptr
	messageBox               uintptr
	postQuitMessage          uintptr
	getActiveWindow          uintptr
	getWindowLong            uintptr
	setWindowLong            uintptr
	getWindowLongPtr         uintptr
	setWindowLongPtr         uintptr
	updateWindow             uintptr
	showWindow               uintptr
	registerClassEx          uintptr
	unregisterClass          uintptr
	defWindowProc            uintptr
	getModuleHandle          uintptr
	loadIcon                 uintptr
	loadCursor               uintptr
	createWindowEx           uintptr
	translateMessage         uintptr
	dispatchMessage          uintptr
	setParent                uintptr
)

//Private Declare Function mciSendStringA Lib "winmm.dll" _
//        (ByVal lpstrCommand As String, ByVal lpstrReturnString As String,
//        ByVal uReturnLength As Integer, ByVal hwndCallback As Integer) As Integer

func init() {
	// Library
	lib = win.MustLoadLibrary("user32.dll")

	// Functions
	destroyWindow = win.MustGetProcAddress(lib, "DestroyWindow")
	sendInput = win.MustGetProcAddress(lib, "SendInput")
	keybd_event = win.MustGetProcAddress(lib, "keybd_event")
	getForegroundWindow = win.MustGetProcAddress(lib, "GetForegroundWindow")
	getWindowText = win.MustGetProcAddress(lib, "GetWindowTextW")
	findWindow = win.MustGetProcAddress(lib, "FindWindowW")
	findWindowEx = win.MustGetProcAddress(lib, "FindWindowExW")
	getWindowTextLength = win.MustGetProcAddress(lib, "GetWindowTextLengthW")
	sendMessage = win.MustGetProcAddress(lib, "SendMessageW")
	postMessage = win.MustGetProcAddress(lib, "PostMessageW")
	getMessage = win.MustGetProcAddress(lib, "GetMessageW")
	translateMessage = win.MustGetProcAddress(lib, "TranslateMessage")
	dispatchMessage = win.MustGetProcAddress(lib, "DispatchMessageW")
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
	registerClassEx = win.MustGetProcAddress(lib, "RegisterClassExW")
	unregisterClass = win.MustGetProcAddress(lib, "UnregisterClassW")
	defWindowProc = win.MustGetProcAddress(lib, "DefWindowProcW")
	loadIcon = win.MustGetProcAddress(lib, "LoadIconW")
	loadCursor = win.MustGetProcAddress(lib, "LoadCursorW")
	createWindowEx = win.MustGetProcAddress(lib, "CreateWindowExW")
	postQuitMessage = win.MustGetProcAddress(lib, "PostQuitMessage")
	setParent = win.MustGetProcAddress(lib, "SetParent")
}

// Predefined icon constants
const (
	IDI_APPLICATION = 32512
	IDI_HAND        = 32513
	IDI_QUESTION    = 32514
	IDI_EXCLAMATION = 32515
	IDI_ASTERISK    = 32516
	IDI_WINLOGO     = 32517
	IDI_SHIELD      = 32518
	IDI_WARNING     = IDI_EXCLAMATION
	IDI_ERROR       = IDI_HAND
	IDI_INFORMATION = IDI_ASTERISK
)

// Predefined cursor constants
const (
	IDC_ARROW       = 32512
	IDC_IBEAM       = 32513
	IDC_WAIT        = 32514
	IDC_CROSS       = 32515
	IDC_UPARROW     = 32516
	IDC_SIZENWSE    = 32642
	IDC_SIZENESW    = 32643
	IDC_SIZEWE      = 32644
	IDC_SIZENS      = 32645
	IDC_SIZEALL     = 32646
	IDC_NO          = 32648
	IDC_HAND        = 32649
	IDC_APPSTARTING = 32650
	IDC_HELP        = 32651
	IDC_ICON        = 32641
	IDC_SIZE        = 32640
)

// Predefined brushes constants
const (
	COLOR_3DDKSHADOW              = 21
	COLOR_3DFACE                  = 15
	COLOR_3DHILIGHT               = 20
	COLOR_3DHIGHLIGHT             = 20
	COLOR_3DLIGHT                 = 22
	COLOR_BTNHILIGHT              = 20
	COLOR_3DSHADOW                = 16
	COLOR_ACTIVEBORDER            = 10
	COLOR_ACTIVECAPTION           = 2
	COLOR_APPWORKSPACE            = 12
	COLOR_BACKGROUND              = 1
	COLOR_DESKTOP                 = 1
	COLOR_BTNFACE                 = 15
	COLOR_BTNHIGHLIGHT            = 20
	COLOR_BTNSHADOW               = 16
	COLOR_BTNTEXT                 = 18
	COLOR_CAPTIONTEXT             = 9
	COLOR_GRAYTEXT                = 17
	COLOR_HIGHLIGHT               = 13
	COLOR_HIGHLIGHTTEXT           = 14
	COLOR_INACTIVEBORDER          = 11
	COLOR_INACTIVECAPTION         = 3
	COLOR_INACTIVECAPTIONTEXT     = 19
	COLOR_INFOBK                  = 24
	COLOR_INFOTEXT                = 23
	COLOR_MENU                    = 4
	COLOR_MENUTEXT                = 7
	COLOR_SCROLLBAR               = 0
	COLOR_WINDOW                  = 5
	COLOR_WINDOWFRAME             = 6
	COLOR_WINDOWTEXT              = 8
	COLOR_HOTLIGHT                = 26
	COLOR_GRADIENTACTIVECAPTION   = 27
	COLOR_GRADIENTINACTIVECAPTION = 28
)

// Window style constants
const (
	WS_OVERLAPPED       = 0X00000000
	WS_POPUP            = 0X80000000
	WS_CHILD            = 0X40000000
	WS_MINIMIZE         = 0X20000000
	WS_VISIBLE          = 0X10000000
	WS_DISABLED         = 0X08000000
	WS_CLIPSIBLINGS     = 0X04000000
	WS_CLIPCHILDREN     = 0X02000000
	WS_MAXIMIZE         = 0X01000000
	WS_CAPTION          = 0X00C00000
	WS_BORDER           = 0X00800000
	WS_DLGFRAME         = 0X00400000
	WS_VSCROLL          = 0X00200000
	WS_HSCROLL          = 0X00100000
	WS_SYSMENU          = 0X00080000
	WS_THICKFRAME       = 0X00040000
	WS_GROUP            = 0X00020000
	WS_TABSTOP          = 0X00010000
	WS_MINIMIZEBOX      = 0X00020000
	WS_MAXIMIZEBOX      = 0X00010000
	WS_TILED            = 0X00000000
	WS_ICONIC           = 0X20000000
	WS_SIZEBOX          = 0X00040000
	WS_OVERLAPPEDWINDOW = 0X00000000 | 0X00C00000 | 0X00080000 | 0X00040000 | 0X00020000 | 0X00010000
	WS_POPUPWINDOW      = 0X80000000 | 0X00800000 | 0X00080000
	WS_CHILDWINDOW      = 0X40000000
)
const CW_USEDEFAULT = ^0x7fffffff

// ComboBox styles
const (
	CBS_SIMPLE            = 0x0001
	CBS_DROPDOWN          = 0x0002
	CBS_DROPDOWNLIST      = 0x0003
	CBS_OWNERDRAWFIXED    = 0x0010
	CBS_OWNERDRAWVARIABLE = 0x0020
	CBS_AUTOHSCROLL       = 0x0040
	CBS_OEMCONVERT        = 0x0080
	CBS_SORT              = 0x0100
	CBS_HASSTRINGS        = 0x0200
	CBS_NOINTEGRALHEIGHT  = 0x0400
	CBS_DISABLENOSCROLL   = 0x0800
	CBS_UPPERCASE         = 0x2000
	CBS_LOWERCASE         = 0x4000
)

const (
	SBS_BOTTOMALIGN             = 0x4
	SBS_HORZ                    = 0x0
	SBS_LEFTALIGN               = 0x2
	SBS_RIGHTALIGN              = 0x4
	SBS_SIZEBOX                 = 0x8
	SBS_SIZEBOXBOTTOMRIGHTALIGN = 0x4
	SBS_SIZEBOXTOPLEFTALIGN     = 0x2
	SBS_SIZEGRIP                = 0x10
	SBS_TOPALIGN                = 0x2
	SBS_VERT                    = 0x1
)

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

// ListBox style
const (
	LBS_NOTIFY            = 0x0001
	LBS_SORT              = 0x0002
	LBS_NOREDRAW          = 0x0004
	LBS_MULTIPLESEL       = 0x0008
	LBS_OWNERDRAWFIXED    = 0x0010
	LBS_OWNERDRAWVARIABLE = 0x0020
	LBS_HASSTRINGS        = 0x0040
	LBS_USETABSTOPS       = 0x0080
	LBS_NOINTEGRALHEIGHT  = 0x0100
	LBS_MULTICOLUMN       = 0x0200
	LBS_WANTKEYBOARDINPUT = 0x0400
	LBS_EXTENDEDSEL       = 0x0800
	LBS_DISABLENOSCROLL   = 0x1000
	LBS_NODATA            = 0x2000
	LBS_NOSEL             = 0x4000
	LBS_COMBOBOX          = 0x8000
	LBS_STANDARD          = LBS_NOTIFY | LBS_SORT | WS_BORDER | WS_VSCROLL
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
func UpdateWindow(hwnd win.HWND) bool {
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

func ShowWindow(hWnd win.HWND, nCmdShow int32) bool {
	ret, _, _ := syscall.Syscall(showWindow, 2,
		uintptr(hWnd),
		uintptr(nCmdShow),
		0)

	return ret != 0
}

// Window class styles
const (
	CS_VREDRAW         = 0x00000001
	CS_HREDRAW         = 0x00000002
	CS_KEYCVTWINDOW    = 0x00000004
	CS_DBLCLKS         = 0x00000008
	CS_OWNDC           = 0x00000020
	CS_CLASSDC         = 0x00000040
	CS_PARENTDC        = 0x00000080
	CS_NOKEYCVT        = 0x00000100
	CS_NOCLOSE         = 0x00000200
	CS_SAVEBITS        = 0x00000800
	CS_BYTEALIGNCLIENT = 0x00001000
	CS_BYTEALIGNWINDOW = 0x00002000
	CS_GLOBALCLASS     = 0x00004000
	CS_IME             = 0x00010000
	CS_DROPSHADOW      = 0x00020000
)

// Button style constants
const (
	BS_3STATE          = 5
	BS_AUTO3STATE      = 6
	BS_AUTOCHECKBOX    = 3
	BS_AUTORADIOBUTTON = 9
	BS_BITMAP          = 128
	BS_BOTTOM          = 0X800
	BS_CENTER          = 0X300
	BS_CHECKBOX        = 2
	BS_DEFPUSHBUTTON   = 1
	BS_GROUPBOX        = 7
	BS_ICON            = 64
	BS_LEFT            = 256
	BS_LEFTTEXT        = 32
	BS_MULTILINE       = 0X2000
	BS_NOTIFY          = 0X4000
	BS_OWNERDRAW       = 0XB
	BS_PUSHBUTTON      = 0
	BS_PUSHLIKE        = 4096
	BS_RADIOBUTTON     = 4
	BS_RIGHT           = 512
	BS_RIGHTBUTTON     = 32
	BS_SPLITBUTTON     = 0x0000000c
	BS_TEXT            = 0
	BS_TOP             = 0X400
	BS_USERBUTTON      = 8
	BS_VCENTER         = 0XC00
	BS_FLAT            = 0X8000
)

type WNDCLASSEX struct {
	CbSize        uint32
	Style         uint32
	LpfnWndProc   uintptr
	CbClsExtra    int32
	CbWndExtra    int32
	HInstance     win.HINSTANCE
	HIcon         win.HICON
	HCursor       win.HCURSOR
	HbrBackground win.HBRUSH
	LpszMenuName  *uint16
	LpszClassName *uint16
	HIconSm       win.HICON
}

func RegisterClassEx(windowClass *WNDCLASSEX) win.ATOM {
	ret, _, _ := syscall.Syscall(registerClassEx, 1,
		uintptr(unsafe.Pointer(windowClass)),
		0,
		0)

	return win.ATOM(ret)
}
func UnregisterClass(name *uint16) bool {
	ret, _, _ := syscall.Syscall(unregisterClass, 1,
		uintptr(unsafe.Pointer(name)),
		0,
		0)

	return ret != 0
}

func DefWindowProc(hWnd win.HWND, Msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := syscall.Syscall6(defWindowProc, 4,
		uintptr(hWnd),
		uintptr(Msg),
		wParam,
		lParam,
		0,
		0)

	return ret
}

func LoadIcon(hInstance win.HINSTANCE, lpIconName *uint16) win.HICON {
	ret, _, _ := syscall.Syscall(loadIcon, 2,
		uintptr(hInstance),
		uintptr(unsafe.Pointer(lpIconName)),
		0)

	return win.HICON(ret)
}

func LoadCursor(hInstance win.HINSTANCE, lpCursorName *uint16) win.HCURSOR {
	ret, _, _ := syscall.Syscall(loadCursor, 2,
		uintptr(hInstance),
		uintptr(unsafe.Pointer(lpCursorName)),
		0)

	return win.HCURSOR(ret)
}

func CreateWindowEx(dwExStyle uint32, lpClassName, lpWindowName *uint16, dwStyle uint32, x, y, nWidth, nHeight int32, hWndParent win.HWND, hMenu win.HMENU, hInstance win.HINSTANCE, lpParam unsafe.Pointer) win.HWND {
	ret, _, _ := syscall.Syscall12(createWindowEx, 12,
		uintptr(dwExStyle),
		uintptr(unsafe.Pointer(lpClassName)),
		uintptr(unsafe.Pointer(lpWindowName)),
		uintptr(dwStyle),
		uintptr(x),
		uintptr(y),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(hWndParent),
		uintptr(hMenu),
		uintptr(hInstance),
		uintptr(lpParam))

	return win.HWND(ret)
}

func SetParent(hWnd win.HWND, parentHWnd win.HWND) win.HWND {
	ret, _, _ := syscall.Syscall(setParent, 2,
		uintptr(hWnd),
		uintptr(parentHWnd),
		0)

	return win.HWND(ret)
}

type POINT struct {
	X, Y int32
}
type MSG struct {
	HWnd    win.HWND
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      POINT
}

func GetMessage(msg *MSG, hWnd win.HWND, msgFilterMin, msgFilterMax uint32) int32 {
	ret, _, _ := syscall.Syscall6(getMessage, 4,
		uintptr(unsafe.Pointer(msg)),
		uintptr(hWnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax),
		0,
		0)

	return int32(ret)
}
func TranslateMessage(msg *MSG) bool {
	ret, _, _ := syscall.Syscall(translateMessage, 1,
		uintptr(unsafe.Pointer(msg)),
		0,
		0)

	return ret != 0
}
func DispatchMessage(msg *MSG) uintptr {
	ret, _, _ := syscall.Syscall(dispatchMessage, 1,
		uintptr(unsafe.Pointer(msg)),
		0,
		0)

	return ret
}

func DestroyWindow(hWnd win.HWND) bool {
	ret, _, _ := syscall.Syscall(destroyWindow, 1,
		uintptr(hWnd),
		0,
		0)

	return ret != 0
}

func PostQuitMessage(exitCode int32) {
	syscall.Syscall(postQuitMessage, 1,
		uintptr(exitCode),
		0,
		0)
}

func Free() {
	syscall.FreeLibrary(syscall.Handle(lib))
}
