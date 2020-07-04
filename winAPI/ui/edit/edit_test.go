package edit

import (
	"testing"

	"github.com/xiongyejun/xyjgo/winAPI/UI/form"
)

func Test_func(t *testing.T) {
	f := form.New()
	// f.AddClickFunc(func() {
	// 	for i := range f.Controls {
	// 		fmt.Println(i, f.Controls[i].GetHwnd(), f.Controls[i].GetName())
	// 	}
	// })

	f.Create(nil)

	f.AddControl(New())

	f.Show()
	f.LoopMessage()

	t.Log(1)
}
