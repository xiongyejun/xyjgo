package user32

import (
	"testing"
	"time"
)

func Test_func(t *testing.T) {
	defer Free()

	var i int = 0
	for i < 10 {
		time.Sleep(2 * time.Second)
		hwnd := GetActiveWindow()
		t.Log(hwnd)
		t.Log(GetWindowText(hwnd))
		i++
	}

}
