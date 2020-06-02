// 窗体
// 1、注册窗体类
// 2、创建窗体
// 3、显示窗体
// 4、轮询消息
// 5、销毁窗体

package form

import (
	"errors"
	"syscall"
	"unsafe"

	"github.com/xiongyejun/xyjgo/winAPI/ui"

	"github.com/xiongyejun/xyjgo/winAPI/kernel32"
	"github.com/xiongyejun/xyjgo/winAPI/user32"
	"github.com/xiongyejun/xyjgo/winAPI/win"
)

const (
	CLASS_NAME string = "xyj"
)

type Form struct {
	hwnd win.HWND

	lpszClassName *uint16
	handle        win.HANDLE

	name   string
	lpName *uint16

	Left, Top, Width, Height int32

	Controls []ui.Controler
}

var handle win.HANDLE = kernel32.GetModuleHandle(nil)
var lpszClassName *uint16 = syscall.StringToUTF16Ptr(CLASS_NAME)
var fhwnd win.HWND

func init() {
	if err := registerWindowClass(); err != nil {
		panic(err)
	}
}

func New() (ret *Form) {
	ret = new(Form)
	ret.SetName("xyjForm")
	ret.lpszClassName = lpszClassName
	ret.handle = handle
	ret.Left = user32.CW_USEDEFAULT
	ret.Top = user32.CW_USEDEFAULT
	ret.Width = user32.CW_USEDEFAULT
	ret.Height = user32.CW_USEDEFAULT

	return
}

// 创建窗体
func (me *Form) Create(parent ui.Container) {
	me.hwnd = win.HWND(user32.CreateWindowEx(0, me.lpszClassName, me.lpName, user32.WS_OVERLAPPEDWINDOW, me.Left, me.Top, me.Width, me.Height, 0, 0, me.handle, nil))
	if me.hwnd == 0 {
		panic("CreateWindowEx == 0")
	}
	fhwnd = me.hwnd
}

// 显示窗体
func (me *Form) Show() {
	user32.ShowWindow(me.hwnd, user32.SW_SHOWNORMAL)
	user32.UpdateWindow(me.hwnd)

}

// 进入消息循环
func (me *Form) LoopMessage() {
	var msg *user32.MSG = new(user32.MSG)
	// 返回值：如果函数取得WM_QUIT之外的其他消息，返回非零值
	// 如果函数取得WM_QUIT消息，返回值是零
	// 如果出现了错误，返回值是-1
	for user32.GetMessage(msg, 0, 0, 0) > 0 {
		user32.TranslateMessage(msg)
		user32.DispatchMessage(msg)
	}
}

// 获取子控件
func (me *Form) GetControls() []ui.Controler {
	return me.Controls
}
func (me *Form) AddControl(ctl ui.Controler) {
	ctl.Create(me)
	me.Controls = append(me.Controls, ctl)
}

// 设置窗体名称
func (me *Form) SetName(name string) {
	me.name = name

	me.lpName = syscall.StringToUTF16Ptr(me.name)
}

func (me *Form) GetHwnd() uintptr {
	return me.hwnd
}
func (me *Form) GetHandle() uintptr {
	return me.handle
}

// 回调函数
func wndProc(hWnd win.HWND, uMsg uint32, wParam, lParam uintptr) uintptr {
	switch uMsg {
	case user32.WM_DESTROY:
		user32.DestroyWindow(fhwnd)
		user32.PostQuitMessage(0)
		user32.UnregisterClass(lpszClassName)

		// case user32.WM_CREATE:
		// 	user32.CreateWindowEx(0, syscall.StringToUTF16Ptr("BUTTON"), syscall.StringToUTF16Ptr("BUTTON"), user32.WS_CHILD|user32.WS_VISIBLE, 10, 10, 80, 20, hWnd, 0, handle, nil)

	}
	return user32.DefWindowProc(hWnd, uMsg, wParam, lParam)
}

// 注册窗体类
func registerWindowClass() (err error) {
	wc := user32.WNDCLASSEX{}
	wc.CbSize = uint32(unsafe.Sizeof(wc))
	wc.Style = user32.CS_HREDRAW | user32.CS_VREDRAW
	wc.LpfnWndProc = syscall.NewCallback(wndProc)
	wc.HInstance = handle
	wc.HIcon = user32.LoadIcon(0, win.MAKEINTRESOURCE(user32.IDI_APPLICATION))
	wc.HCursor = user32.LoadCursor(0, win.MAKEINTRESOURCE(user32.IDC_ARROW))
	wc.HbrBackground = user32.COLOR_WINDOW
	wc.LpszClassName = lpszClassName

	ret := user32.RegisterClassEx(&wc)
	if ret == 0 {
		err = errors.New("RegisterClassEx == 0")
		return
	}

	return
}
