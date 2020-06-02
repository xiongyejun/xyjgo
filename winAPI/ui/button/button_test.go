package button

import (
	"strconv"
	"testing"

	"github.com/xiongyejun/xyjgo/winAPI/UI/form"
)

func Test_func(t *testing.T) {
	f := form.New()
	f.Create(nil)

	for i := 0; i < 3; i++ {
		b := New()
		b.Top = int32(40*i + 5)
		b.SetName(strconv.Itoa(i))
		b.Create(f)
		f.AddControl(b)
	}

	f.Show()
	f.LoopMessage()

	t.Log(1)
}
