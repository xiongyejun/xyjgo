package pressKey

import (
	"testing"
)

func Test_func(t *testing.T) {
	k := New()
	k.Add('A', 3000, 0)
	k.Add('T', 3000, 0)

	k.Start()
}
