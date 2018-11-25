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

func init() {
	tsl = translate.NewYouDao()
}

func main() {
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
		tsl = translate.NewBaiDu()
	case "youdao":
		tsl = translate.NewYouDao()

	default:
		if ret, err := tsl.Translate(strings.Join(tokens, " ")); err != nil {
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
