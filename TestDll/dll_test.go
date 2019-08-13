package main

import (
	"encoding/binary"
	"testing"
	"unsafe"
)

func Test_func(t *testing.T) {
	var i int32 = 10
	pi := uintptr(unsafe.Pointer(&i))

	t.Log("int32 size = ", binary.Size(i), "uintptr = ", binary.Size(pi))

	var str string = "str"
	pstr := &str
	t.Log("str=", str, "pstr=", pstr)

	ppstr := uintptr(unsafe.Pointer(pstr))

	t.Log(pstr, ppstr)
	//	var pint *int = (*int)(ppstr)
	//	t.Log(pint)
}
