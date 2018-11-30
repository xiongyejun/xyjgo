package keyboard

import (
	"testing"
	"time"
)

func Test_press(t *testing.T) {
	time.Sleep(1 * time.Second)

	defer Free()
	t.Log(Press(VK_A))
	t.Log(Press(VK_G))
	t.Log(Press(VK_S))
	t.Log(Press(VK_1))
	t.Log(Press(VK_9))

}
