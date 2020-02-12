package main

import (
	"fmt"
)

func main() {

	fmt.Println("ok")
}

func printHelp() {
	fmt.Printf(`
	mygo xs	<set.txt>	下载小说
	mygo epub	创建epub(iphone iBooks)
	`)
}
