package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/xiongyejun/xyjgo/rePolish"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("gocalc <中缀表达式>")
		return
	}

	var str = os.Args[1]
	str = strings.Replace(str, " ", "", -1)
	var arr []string
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			var j int = i + 1
			for j < len(str) && (str[j] >= '0' && str[j] <= '9') {
				j++
			}
			arr = append(arr, string(str[i:j]))
			i = j - 1
		} else {
			arr = append(arr, string(str[i]))
		}
	}

	fmt.Println("中缀表达式:", arr)
	var err error
	if arr, err = rePolish.RePolish(arr); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("后缀表达式:", arr)

	var f float64
	if f, err = rePolish.Calc(arr); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("计算结果为: %f\n", f)
}
