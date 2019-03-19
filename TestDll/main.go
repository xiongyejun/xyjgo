package main

import (
	"C"
	"fmt"
)

func main() {

}

//Hello :
//export Hello
func Hello() {
	fmt.Println("hello from golang")
}

//Sum :
//export Sum
func Sum(a, b int) int {
	return a + b
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
