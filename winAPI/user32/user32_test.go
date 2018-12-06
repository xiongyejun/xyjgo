package user32

import (
	"testing"
	"time"
)

func Test_func(t *testing.T) {
	t.Log("屏蔽开始")
	BlockInput(1)

	time.Sleep(3 * time.Second)

	BlockInput(0)

	t.Log("屏蔽结束")
}
