package main

import (
	"fmt"
	"time"
)

func main() {
	var i int32 = 12
	var a int32 = 65
	var str string = "string"

	fmt.Println("address i =", &i)
	fmt.Println("address a =", &a)
	fmt.Println("address str =", &str)
	println(str)

	time.Sleep(1e9 * 9999999)
}
