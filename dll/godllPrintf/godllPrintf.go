package main

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/xiongyejun/xyjgo/vbatd"
)

func main() {
}

//export Sprintf
func Sprintf(pformat, pParamArray, nCount int) (ptr unsafe.Pointer, lenth int) {
	var err error
	var str string

	var it interface{}
	if it, err = vbatd.Variant2interface(uintptr(pformat)); err != nil {
		str = err.Error()
		return str2ptr(str)
	}

	var strformat string
	var ok bool
	if strformat, ok = it.(string); !ok {
		return str2ptr("pformat 指向的不是VBA Variant的String.")
	}

	var its []interface{}
	if its, err = vbatd.Variants2interfaces(uintptr(pParamArray), nCount); err != nil {
		str = err.Error()
		return str2ptr(str)
	}

	str = fmt.Sprintf(strformat, its...)

	return str2ptr(str)
}

//export Sum
func Sum(a, b int) int {
	return a + b
}

//export Free
func Free(p unsafe.Pointer) {
	C.free(p)
}

func str2ptr(str string) (ptr unsafe.Pointer, lenth int) {
	b := []byte(str)
	lenth = len(b)
	ptr = C.malloc(C.uint(lenth))
	C.memcpy(ptr, unsafe.Pointer(&b[0]), C.uint(lenth))

	return
}
