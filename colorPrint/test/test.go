package main

import (
	"fmt"

	"github.com/xiongyejun/xyjgo/colorPrint"
)

func main() {
	colorPrint.SetColor(colorPrint.Red, colorPrint.Black)
	fmt.Println("test red")

	colorPrint.SetColor(colorPrint.Blue, colorPrint.Black)
	fmt.Println("test Blue")

	colorPrint.SetColor(colorPrint.Green, colorPrint.Black)
	fmt.Println("test Green")

	colorPrint.ReSetColor()
}
