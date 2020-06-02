package button

import (
	"syscall"

	// win1 "github.com/lxn/win"
	"github.com/xiongyejun/xyjgo/winAPI/ui"
	"github.com/xiongyejun/xyjgo/winAPI/user32"
	"github.com/xiongyejun/xyjgo/winAPI/win"
)

const (
	CLASS_NAME string = "BUTTON"
)

type Button struct {
	hwnd win.HWND

	lpszClassName *uint16

	name   string
	lpName *uint16

	Left, Top, Width, Height int32
}

func init() {

}

func New() (ret *Button) {
	ret = new(Button)
	ret.SetName("xyjButton")
	ret.lpszClassName = syscall.StringToUTF16Ptr(CLASS_NAME)
	ret.Left = 5
	ret.Width = 80
	ret.Top = 10
	ret.Height = 30

	return
}

// 显示
func (me *Button) Create(parent ui.Container) {
	me.hwnd = win.HWND(user32.CreateWindowEx(0, me.lpszClassName, me.lpName, user32.WS_CHILD|user32.WS_VISIBLE, me.Left, me.Top, me.Width, me.Height, parent.GetHwnd(), uintptr(len(parent.GetControls())+1), parent.GetHandle(), nil))
	if me.hwnd == 0 {
		panic("CreateWindowEx == 0")
	}
	// user32.SetParent(me.hwnd, parent)
}

func (me *Button) GetHwnd() uintptr {
	return me.hwnd
}

// 设置名称
func (me *Button) SetName(name string) {
	me.name = name

	me.lpName = syscall.StringToUTF16Ptr(me.name)
}
