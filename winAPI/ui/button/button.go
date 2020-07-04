package button

import (
	"github.com/xiongyejun/xyjgo/winAPI/ui"
)

const (
	CLASS_NAME string = "BUTTON"
)

type Button struct {
	*ui.Control
}

func init() {

}

func New() (ret *Button) {
	ret = new(Button)
	ret.Control = ui.StdControl[ui.BUTTON]
	ret.SetName("xyjButton")
	ret.Left = 5
	ret.Width = 80
	ret.Top = 10
	ret.Height = 30

	return
}
