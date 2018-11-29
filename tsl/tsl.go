// 翻译

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/xiongyejun/xyjgo/colorPrint"
	"github.com/xiongyejun/xyjgo/translate"
)

var tsl translate.ITranslate

func main() {
	var err error
	if tsl, err = translate.NewYouDao(); err != nil {
		fmt.Println(err)
		return
	}

	printCmd()
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter Cmd->")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)
		if line == "q" || line == "e" {
			break
		}
		tokens := strings.Split(line, " ")
		printCmd()
		handleCommands(tokens)
	}

}

func handleCommands(tokens []string) {
	switch tokens[0] {
	case "printCmd":
		printCmd()
	case "baidu":
		var err error
		if tsl, err = translate.NewBaiDu(); err != nil {
			fmt.Println(err)
			return
		}
	case "youdao":
		var err error
		if tsl, err = translate.NewYouDao(); err != nil {
			fmt.Println(err)
			return
		}

	default:
		if ret, err := tsl.Translate(strings.Join(tokens, " "), true); err != nil {
			fmt.Println(err)
		} else {
			colorPrint.SetColor(colorPrint.White, colorPrint.DarkMagenta)
			fmt.Println(ret)
			colorPrint.ReSetColor()
		}
	}
}

func printCmd() {
	colorPrint.SetColor(colorPrint.Green, colorPrint.Black)

	fmt.Println(` Enter following commands to control:
	
 printCmd -- 查看命令
 baidu -- baidu翻译
 youdao -- youdao翻译(默认)
`)
	colorPrint.ReSetColor()
}
