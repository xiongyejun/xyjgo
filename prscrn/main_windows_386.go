// 调用微信截图dll

package main

import (
	"fmt"
	"os"
	"syscall"
)

// 需先设置环境为32位的
func main() {
	dllPath := os.Getenv("GOPATH") + `\src\github.com\xiongyejun\xyjgo\prscrn\PrScrn.dll`
	if dll, err := syscall.LoadLibrary(dllPath); err != nil {
		fmt.Println(err.Error() + " LoadLibrary")
	} else {
		defer syscall.FreeLibrary(dll)
		if ps, err := syscall.GetProcAddress(dll, "PrScrn"); err != nil {
			fmt.Println(err.Error() + " GetProcAddress")
		} else {

			syscall.Syscall(ps, 0, 0, 0, 0)
		}
	}
}
