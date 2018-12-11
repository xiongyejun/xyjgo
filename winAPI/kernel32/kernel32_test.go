package kernel32

import (
	"fmt"
	"syscall"
	"testing"
	"unsafe"
)

func Test_shortPath(t *testing.T) {
	var shortBufferSize uint32 = 512
	var shortPath = make([]uint16, shortBufferSize)
	ret := GetShortPathName(`C:\Users\Administrator\Music\Amani.mp3`, uintptr(unsafe.Pointer(&shortPath[0])), shortBufferSize)
	fmt.Println(ret)
	print("shortpath=", syscall.UTF16ToString(shortPath))
}
