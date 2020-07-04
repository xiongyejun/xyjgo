package ui

import (
	"syscall"

	// "github.com/lxn/win"
	"github.com/xiongyejun/xyjgo/winAPI/user32"
)

type Container interface {
	GetControls() []Controler
	GetHandle() uintptr

	Controler
}

type Controler interface {
	Create(Container)
	GetHwnd() uintptr
	GetName() string
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

	STD_CONTROL_COUNT
)

type Control struct {
	Name      string
	ClassName string
	Style     uint32

	Left, Top, Width, Height int32
	Parent                   Controler
	hwnd                     uintptr
	id                       int
}

func (me *Control) Create(parent Container) {
	me.id = len(parent.GetControls()) + 1
	me.Parent = parent

	me.hwnd = user32.CreateWindowEx(0, syscall.StringToUTF16Ptr(me.ClassName), syscall.StringToUTF16Ptr(me.Name), me.Style, me.Left, me.Top, me.Width, me.Height, parent.GetHwnd(), uintptr(me.id), parent.GetHandle(), nil)
	if me.hwnd == 0 {
		panic("CreateWindowEx == 0")
	}
}

func (me *Control) GetHwnd() uintptr {
	return me.hwnd
}

// 设置名称
func (me *Control) SetName(name string) {
	me.Name = name
}
func (me *Control) GetName() string {
	return me.Name
}

// windows 标准控件
var StdControl []*Control = make([]*Control, STD_CONTROL_COUNT)

func init() {
	StdControl[STATIC] = &Control{
		Name:      "Static",
		ClassName: "STATIC",
		Style:     user32.WS_CHILD | user32.WS_VISIBLE,
		Left:      10,
		Top:       10,
		Width:     80,
		Height:    20,
	}

	StdControl[BUTTON] = &Control{
		Name:      "Button",
		ClassName: "BUTTON",
		Style:     user32.WS_CHILD | user32.WS_VISIBLE,
		Left:      10,
		Top:       10,
		Width:     80,
		Height:    20,
	}

	StdControl[EDIT] = &Control{
		ClassName: "Edit",
		Style:     user32.WS_CHILD | user32.WS_VISIBLE | user32.WS_BORDER | user32.ES_AUTOHSCROLL | user32.ES_AUTOVSCROLL,
		Left:      10,
		Top:       10,
		Width:     80,
		Height:    20,
	}

	StdControl[GROUP_BOX] = &Control{
		Name:      "GroupBox",
		ClassName: "BUTTON",
		Style:     user32.WS_CHILD | user32.WS_VISIBLE | user32.BS_GROUPBOX,
		Left:      10,
		Top:       10,
		Width:     130,
		Height:    200,
	}

	StdControl[RADIO_BUTTON] = &Control{
		Name:      "RadioButton",
		ClassName: "BUTTON",
		Style:     user32.WS_CHILD | user32.WS_VISIBLE | user32.BS_AUTORADIOBUTTON,
		Left:      10,
		Top:       10,
		Width:     120,
		Height:    20,
	}

	StdControl[CHECK_BOX] = &Control{
		Name:      "CheckBox",
		ClassName: "BUTTON",
		Style:     user32.WS_CHILD | user32.WS_VISIBLE | user32.BS_AUTOCHECKBOX,
		Left:      10,
		Top:       10,
		Width:     120,
		Height:    20,
	}

	StdControl[LIST_BOX] = &Control{
		Name:      "ListBbox",
		ClassName: "LISTBOX",
		Style:     user32.WS_CHILD | user32.WS_VISIBLE | user32.LBS_STANDARD,
		Left:      10,
		Top:       10,
		Width:     80,
		Height:    90,
	}

	StdControl[COMBO_BOX] = &Control{
		ClassName: "COMBOBOX",
		Style:     user32.WS_CHILD | user32.WS_VISIBLE | user32.WS_VSCROLL | user32.CBS_AUTOHSCROLL | user32.CBS_DROPDOWNLIST,
		Left:      10,
		Top:       10,
		Width:     100,
		Height:    90,
	}

	StdControl[SCROLLBAR] = &Control{
		ClassName: "SCROLLBAR",
		Style:     user32.WS_CHILD | user32.WS_VISIBLE | user32.SBS_HORZ,
		Left:      10,
		Top:       10,
		Width:     200,
		Height:    20,
	}
}
