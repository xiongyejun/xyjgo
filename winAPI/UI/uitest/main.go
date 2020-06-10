package main

import (
	"github.com/xiongyejun/xyjgo/winAPI/ui"
	"github.com/xiongyejun/xyjgo/winAPI/ui/form"
)

func main() {
	f := form.New()
	f.Create(nil)

	var top int32 = 5
	for i := 0; i < ui.STD_CONTROL_COUNT; i++ {
		b := ui.StdControl[i]
		b.Top = top
		f.AddControl(b)
		top += b.Height
		top += 5
	}

	f.Show()
	f.LoopMessage()
}
