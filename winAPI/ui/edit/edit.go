package edit

import (
	"syscall"
	"unsafe"

	"github.com/xiongyejun/xyjgo/winAPI/ui"
	"github.com/xiongyejun/xyjgo/winAPI/user32"
)

const (
	CLASS_NAME string = "BUTTON"
)

type Edit struct {
	*ui.Control
	text string
}

func init() {

}

func New() (ret *Edit) {
	ret = new(Edit)
	ret.Control = ui.StdControl[ui.EDIT]
	ret.Name = "xyjEdit"
	ret.Left = 5
	ret.Width = 80
	ret.Top = 10
	ret.Height = 30

	return
}

//设置Edit文本最大长度
func (me *Edit) SetMaxLength(size uintptr /*字节长度*/) {

}

func (me *Edit) GetText() (ret string) {
	return me.text
}
func (me *Edit) SetText(txt string) {
	user32.SendMessage(me.GetHwnd(), user32.WM_SETTEXT, 0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(txt))))

	me.text = txt
}
