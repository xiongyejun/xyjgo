package form

import (
	"testing"
)

func Test_func(t *testing.T) {
	f := New()
	f.Create(nil)
	f.Show()
	f.LoopMessage()

	t.Log(1)
}