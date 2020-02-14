// 翻译
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/xiongyejun/xyjgo/colorPrint"
	"github.com/xiongyejun/xyjgo/translate/baidu"
	"github.com/xiongyejun/xyjgo/translate/google"
	"github.com/xiongyejun/xyjgo/translate/youdao"
)

type ITranslate interface {
	Translate(value string) (ret string, err error)
	Speak(value string) (err error)
}

var tsl ITranslate
var strTSL string = "google"

func main() {
	var err error
	if tsl, err = google.New(); err != nil {
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
		handleCommands(tokens)
	}

}

func handleCommands(tokens []string) {
	switch tokens[0] {
	case "help":
		printCmd()
	case "baidu":
		var err error
		if tsl, err = baidu.New(); err != nil {
			fmt.Println(err)
			return
		}
		strTSL = "baidu"
	case "youdao":
		var err error
		if tsl, err = youdao.New(); err != nil {
			fmt.Println(err)
			return
		}
		strTSL = "youdao"
	case "google":
		var err error
		if tsl, err = google.New(); err != nil {
			fmt.Println(err)
			return
		}
		strTSL = "google"

	default:
		if ret, err := tsl.Translate(strings.Join(tokens, " ")); err != nil {
			fmt.Println(err)
		} else {
			colorPrint.SetColor(colorPrint.White, colorPrint.DarkMagenta)
			fmt.Println(strTSL, "翻译：")
			fmt.Println(ret)
			colorPrint.ReSetColor()
		}
	}
}

func printCmd() {
	colorPrint.SetColor(colorPrint.Green, colorPrint.Black)

	fmt.Println(` Enter following commands to control:
	
 help -- 查看命令
 baidu -- baidu翻译
 youdao -- youdao翻译
 google -- google翻译(默认)
`)
	colorPrint.ReSetColor()
}
