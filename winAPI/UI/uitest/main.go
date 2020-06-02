package main

import (
	"github.com/xiongyejun/xyjgo/winAPI/ui"
	"github.com/xiongyejun/xyjgo/winAPI/ui/form"
)

func main() {
	f := form.New()
	f.Create(nil)

	var top int32 = 5
	for i := 0; i < ui.CONTROL_COUNT; i++ {
		b := ui.DefaultControl[i]
		b.Top = top
		f.AddControl(b)
		top += b.Height
		top += 5
	}

	f.Show()
	f.LoopMessage()
}
