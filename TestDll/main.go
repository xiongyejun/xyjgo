package main

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"
import (
	"unsafe"
)

func main() {
}

//export GetStr
func GetStr() string {
	return "Go string测试。haha"
}

//export ReurnSlice
func ReurnSlice() (ptr unsafe.Pointer, iLen, iCap int) {
	s := make([]int, 0, 10)
	s = append(s, 23)
	s = append(s, 231)
	s = append(s, 234)

	ptr = C.malloc(C.uint(4 * cap(s)))
	C.memcpy(ptr, unsafe.Pointer(&s[0]), C.uint(4*cap(s)))
	return ptr, len(s), cap(s)
}

//export Free
func Free(p unsafe.Pointer) {
	C.free(p)
}

//export Sum
func Sum(a, b int) int {
	return a + b
}

// 原文：https://studygolang.com/articles/19000#reply0
// 编译命令： go build --buildmode=c-shared -o go.dll main.go

// 另一种方法 https://blog.csdn.net/qq_30549833/article/details/86157744
//EXPORTS
//	Sum
//	Hello

// 3、gcc yy.def xx.a -shared -lwinmm -lWs2_32 -o zz.dll -Wl,--out-implib,zz.lib

// vba调用dll遵循的是stdcall
// 1、go build -v -x -buildmode=c-archive -o xx.a
//		单独建立文件夹“c”的目的是：编写了.c文件，go build不指定文件的时候，会把.c也包括进去
//		生成xx.a和xx.h 2个文件
// 2、编写yy.def文件
// 		生成.a和.h文件之后，编写一个C文件，需要导出的函数都用C语言实现，一定要加__stdcall
// 		编写.def, def文件导出名称和C文件的保持一致
// 3、再编译dll
//		gcc.exe c\stdcall.c c\go.def c\go.a -shared -lwinmm -lWs2_32 -o go.dll -Wl,--enable-stdcall-fixup,--out-implib,go.lib

// 内存方面
// https://zhuanlan.zhihu.com/p/46721768
// Golang 是一个有垃圾回收的语言
// 所以任何 Golang 的对象指针不应该直接传给 C 语言代码
// 因为 C 语言可以直接操作内存，会导致 Golang 的运行环境被破坏，产生不可预知的问题
// 所以go的对象指针先用C的malloc申请内存ptr，再memcpy过去，函数返回ptr
// vba调用后，记得要free
