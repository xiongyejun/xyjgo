package user32

import (
	"testing"
)

func Test_func(t *testing.T) {
	t.Logf("%x %x\r\n", 'a', MapVirtualKey('a', 0))
}
