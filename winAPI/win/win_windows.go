package win

import (
	"syscall"
	"unsafe"

	"github.com/axgle/mahonia"
)

const (
	CODE_UTF16 int = iota
	CODE_GBK
)

// 读取dll
func MustLoadLibrary(name string) uintptr {
	lib, err := syscall.LoadLibrary(name)
	if err != nil {
		panic(err)
	}

	return uintptr(lib)
}

// 读取函数
func MustGetProcAddress(lib uintptr, name string) uintptr {
	addr, err := syscall.GetProcAddress(syscall.Handle(lib), name)
	if err != nil {
		panic(err)
	}

	return uintptr(addr)
}

// string指针
func StrPtr(str string, code int) uintptr {
	if str == "" {
		return uintptr(0)
	}

	if code == CODE_UTF16 {
		return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str)))
	}

	if code == CODE_GBK {
		encoder := mahonia.NewEncoder("gbk")
		str = encoder.ConvertString(str)
		b := []byte(str)
		return uintptr(unsafe.Pointer(&b[0]))
	}

	return uintptr(0)
}
