package main

import (
	"C"
	"fmt"
	"unsafe"
)

func main() {

}

//export GetStr
func GetStr() string {
	return "Go string测试。haha"
}

//export Hello
func Hello() {
	fmt.Println("hello from golang")
}

//export OneInt
func OneInt(a int) int {
	return a + 1
}

//export Sum
func Sum(a, b int) int {
	return a + b
}

//export TestPInt
func TestPInt(a, b *int) *int {
	var c int = *a + *b
	return &c
}

//export TestPInt32
func TestPInt32(a, b *int32) *int32 {
	var c int32 = *a + *b
	return &c
}

//export TestPtr
func TestPtr(pa, pb uintptr) uintptr {
	var a int = *(*int)(unsafe.Pointer(pa))
	var b int = *(*int)(unsafe.Pointer(pa))

	var c int = a + b
	return uintptr(unsafe.Pointer(&c))
}

// 原文：https://studygolang.com/articles/19000#reply0
// 编译命令： go build --buildmode=c-shared -o go.dll main.go

// 另一种方法 https://blog.csdn.net/qq_30549833/article/details/86157744
// 1、go build -v -x -buildmode=c-archive -o xx.a
//		生成xx.a和xx.h2个文件
// 2、编写yy.def文件
//EXPORTS
//	Sum
//	Hello

// 3、gcc yy.def xx.a -shared -lwinmm -lWs2_32 -o zz.dll -Wl,--out-implib,zz.lib
