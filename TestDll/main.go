package main

import "C"
import (
	"fmt"
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
// 		生成.a和.h文件之后，编写一个C文件，需要导出的函数都用C语言实现
// 		编写.def, def文件导出名称和C文件的保持一致
// 3、再编译dll
//		gcc.exe c\stdcall.c c\go.def c\go.a -shared -lwinmm -lWs2_32 -o go.dll -Wl,--enable-stdcall-fixup,--out-implib,go.lib
