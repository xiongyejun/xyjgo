package main

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"
import (
	"unsafe"

	"github.com/xiongyejun/xyjgo/ucs2"
)

func str2ptr(str string) (ptr unsafe.Pointer, lenth int) {
	b := []byte(str)
	b, _ = ucs2.FromUTF8(b)

	lenth = len(b)
	ptr = C.malloc(C.size_t(lenth))
	C.memcpy(ptr, unsafe.Pointer(&b[0]), C.size_t(lenth))

	return
}

// 根据go的string，在内存中处理好VBA String的内存，返回StrPtr的地址，这个地址是在VarPtr里保存的值
func str2VarPtr(str string) (ret int32) {
	b := []byte(str)
	b, _ = ucs2.FromUTF8(b)
	// VBA String占10个字节VarPtr 4  +  flag 2（多次打印的猜测）  +  lenth 4
	var lenth int32 = int32(len(b))
	bflag := []byte{0, 0x88}
	blenth := *((*[4]byte)(unsafe.Pointer(&lenth)))
	bflag = append(bflag, blenth[:]...)
	b = append(bflag, b...)
	//	b = append(b, []byte{0, 0, 0, 0}...)

	lenth = int32(len(b))
	ptr := C.malloc(C.size_t(lenth))
	C.memcpy(ptr, unsafe.Pointer(&b[0]), C.size_t(lenth))
	ret = int32(uintptr(ptr))
	ret += 6

	return
}

//export Free
func Free(p unsafe.Pointer) {
	C.free(p)
}
