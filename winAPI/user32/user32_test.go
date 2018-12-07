package user32

import (
	"testing"
	"time"
)

func Test_func(t *testing.T) {
	for {
		time.Sleep(30 * time.Second)
		go func() {
			ret := BlockInput(0)
			t.Log("30s BlockInput loop   ret=", ret)
		}()
	}

}
