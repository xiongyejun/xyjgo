package main

import (
	"fmt"
	"os"

	"github.com/xiongyejun/xyjgo/winAPI/user32"
	"github.com/xiongyejun/xyjgo/winAPI/win"
)

var k int = 0

func main() {
	fmt.Println("enumwindow [-s过滤空白]")
	if len(os.Args) == 1 {
		user32.EnumChildWindows(0, enumProc)
	} else if os.Args[1] == "-s" {
		user32.EnumChildWindows(0, enumProc2)
	}
}

func enumProc(hwnd win.HWND, lParam int) int {
	k++
	fmt.Printf("%d  0x%x  【%s】\n", k, hwnd, user32.GetWindowText(hwnd))

	return 1
}
func enumProc2(hwnd win.HWND, lParam int) int {
	var str string = user32.GetWindowText(hwnd)
	if str != "" {
		k++
		fmt.Printf("%d  0x%x  【%s】\n", k, hwnd, str)
	}

	return 1
}
