package form

import (
	"strconv"
	"testing"

	"github.com/xiongyejun/xyjgo/winAPI/ui/button"
	"github.com/xiongyejun/xyjgo/winAPI/ui/edit"
)

func Test_func(t *testing.T) {
	f := New()

	var count int = 0

	f.Create(nil)

	btn := button.New()
	f.AddControl(btn)

	txt := edit.New()
	txt.Top = btn.Top + btn.Height + 5
	f.AddControl(txt)

	f.AddClickFunc(func() {
		if count%2 == 0 {
			txt.SetText(strconv.Itoa(count))
		} else {
			t.Log(txt.GetText())
		}
		count++
	})

	f.Show()
	f.LoopMessage()

	t.Log(1)
}
