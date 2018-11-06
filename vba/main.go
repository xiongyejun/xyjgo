package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/xiongyejun/xyjgo/colorPrint"
	"github.com/xiongyejun/xyjgo/fileHeader"
	"github.com/xiongyejun/xyjgo/vbaProject"
)

func init() {
	of = new(officeFile)
	of.vp = vbaProject.New()
}

func main() {
	var err error

	if len(os.Args) == 1 {
		fmt.Println("请输入文件名。")
	} else {
		fileName := os.Args[1]

		var f *os.File
		if f, err = os.Open(fileName); err != nil {
			fmt.Println(err)
			return
		}
		var b []byte = make([]byte, 10)
		if _, err = f.Read(b); err != nil {
			fmt.Println(err)
			return
		}
		if fileHeader.IsZIP(b) {
			iof = new(file07)
		} else {
			iof = new(file03)
		}
		of.fileName = fileName

		if of.b, err = iof.readFile(of.fileName); err != nil {
			fmt.Println(err)
		} else {
			if err = of.vp.Parse(of.b); err != nil {
				fmt.Println(err)
			} else {

				r := bufio.NewReader(os.Stdin)
				for {
					printCmd()
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
		}
	}
}

func handleCommands(tokens []string) {
	switch tokens[0] {
	case "ls":
		for i := range of.vp.Module {
			fmt.Println(i, of.vp.Module[i].Name)
		}
		colorPrint.ReSetColor()

	case "show":
		if len(tokens) != 2 {
			fmt.Println(`输入的命令不正确show <ModuleName> -- 查看模块代码`)
			fmt.Printf("%d, %#v\r\n", len(tokens), tokens)
			return
		}

		if str, err := of.vp.GetModuleCode(tokens[1]); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(str)
		}

	case "showi":
		if len(tokens) != 2 {
			fmt.Println(`输入的命令不正确showi <ModuleIndex> -- 查看模块代码`)
			fmt.Printf("%d, %#v\r\n", len(tokens), tokens)
			return
		}
		if index, err := strconv.Atoi(tokens[1]); err != nil {
			fmt.Println(err)
			return
		} else {
			if index > len(of.vp.Module) {
				fmt.Println("越界")
				return
			} else {
				if str, err := of.vp.GetModuleCode(of.vp.Module[index].Name); err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(str)
				}
			}
		}

	default:
		fmt.Println("Unrecognized lib command:", tokens)
	}
}

func printCmd() {
	colorPrint.SetColor(colorPrint.Green, colorPrint.Black)

	fmt.Println(`
 Enter following commands to control:
 ls -- 查看模块列表
 show <ModuleName> -- 查看模块代码
 showi <ModuleIndex> -- 查看模块代码
 e或者q -- 退出
 `)

	colorPrint.ReSetColor()
}
