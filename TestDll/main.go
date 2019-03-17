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
