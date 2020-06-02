package ui

import (
	"syscall"

	"github.com/xiongyejun/xyjgo/winAPI/user32"
)

// "github.com/lxn/win"

type Container interface {
	GetControls() []Controler
	GetHandle() uintptr

	Controler
}

type Controler interface {
	Create(Container)
	GetHwnd() uintptr
}

const (
	STATIC = iota
	BUTTON
	EDIT
	GROUP_BOX
	RADIO_BUTTON
	CHECK_BOX
	LIST_BOX
	COMBO_BOX
	SCROLLBAR

	CONTROL_COUNT
)

type Control struct {
	Name      string
	ClassName string
	Style     uint32

	Left, Top, Width, Height int32
	hwnd                     uintptr
}

func (me *Control) Create(parent Container) {
	me.hwnd = user32.CreateWindowEx(0, syscall.StringToUTF16Ptr(me.ClassName), syscall.StringToUTF16Ptr(me.Name), me.Style, me.Left, me.Top, me.Width, me.Height, parent.GetHwnd(), uintptr(len(parent.GetControls())+1), parent.GetHandle(), nil)
	if me.hwnd == 0 {
		panic("CreateWindowEx == 0")
	}
}

func (me *Control) GetHwnd() uintptr {
	return me.hwnd
}

// windows 标准控件
var DefaultControl []*Control = make([]*Control, CONTROL_COUNT)

func init() {
	DefaultControl[STATIC] = &Control{
		Name:      "STATIC",
		ClassName: "STATIC",
		Style:     user32.WS_CHILD | user32.WS_VISIBLE,
		Left:      10,
		Top:       10,
		Width:     80,
		Height:    20,
	}

	DefaultControl[BUTTON] = &Control{
		Name:      "BUTTON",
		ClassName: "BUTTON",
		Style:     user32.WS_CHILD | user32.WS_VISIBLE,
		Left:      10,
		Top:       10,
		Width:     80,
		Height:    20,
	}

	DefaultControl[EDIT] = &Control{
		ClassName: "EDIT",
		Style:     user32.WS_CHILD | user32.WS_VISIBLE | user32.WS_BORDER,
		Left:      10,
		Top:       10,
		Width:     80,
		Height:    20,
	}

	DefaultControl[GROUP_BOX] = &Control{
		Name:      "GROUP BOX",
		ClassName: "BUTTON",
		Style:     user32.WS_CHILD | user32.WS_VISIBLE | user32.BS_GROUPBOX,
		Left:      10,
		Top:       10,
		Width:     130,
		Height:    200,
	}

	DefaultControl[RADIO_BUTTON] = &Control{
		Name:      "RADIO BUTTON",
		ClassName: "BUTTON",
		Style:     user32.WS_CHILD | user32.WS_VISIBLE | user32.BS_AUTORADIOBUTTON,
		Left:      10,
		Top:       10,
		Width:     120,
		Height:    20,
	}

	DefaultControl[CHECK_BOX] = &Control{
		Name:      "CHECK_BOX",
		ClassName: "BUTTON",
		Style:     user32.WS_CHILD | user32.WS_VISIBLE | user32.BS_AUTOCHECKBOX,
		Left:      10,
		Top:       10,
		Width:     120,
		Height:    20,
	}

	DefaultControl[LIST_BOX] = &Control{
		Name:      "LIST BOX",
		ClassName: "LISTBOX",
		Style:     user32.WS_CHILD | user32.WS_VISIBLE | user32.LBS_STANDARD,
		Left:      10,
		Top:       10,
		Width:     80,
		Height:    90,
	}

	DefaultControl[COMBO_BOX] = &Control{
		ClassName: "COMBOBOX",
		Style:     user32.WS_CHILD | user32.WS_VISIBLE | user32.WS_VSCROLL | user32.CBS_AUTOHSCROLL | user32.CBS_DROPDOWNLIST,
		Left:      10,
		Top:       10,
		Width:     100,
		Height:    90,
	}

	DefaultControl[SCROLLBAR] = &Control{
		ClassName: "SCROLLBAR",
		Style:     user32.WS_CHILD | user32.WS_VISIBLE | user32.SBS_HORZ,
		Left:      10,
		Top:       10,
		Width:     200,
		Height:    20,
	}
}
