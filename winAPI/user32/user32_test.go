package user32

import (
	"testing"
	"time"
)

func Test_func(t *testing.T) {

	for i := 0; i < 10; i++ {
		hwnd := GetForegroundWindow()
		t.Logf("%x %s\r\n", hwnd, GetWindowText(hwnd))
		time.Sleep(time.Second)
	}
}
