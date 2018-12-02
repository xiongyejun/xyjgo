package keyboard

import (
	"testing"
	"time"
)

func Test_press(t *testing.T) {
	time.Sleep(1 * time.Second)

	defer Free()
	t.Log('e')
	t.Logf("%x %x\r\n", arrKeyTable['a'].vk, arrKeyTable['a'].scan)
	Press('e')
	t.Log('f')
	Press('f')

}
