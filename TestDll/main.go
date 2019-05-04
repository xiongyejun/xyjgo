package main

import (
	"C"
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

//export TestPInt
func TestPInt(a, b *int) *int {
	var c int = *a + *b
	return &c
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
