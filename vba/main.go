package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
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
			of.iof = new(file07)
		} else {
			of.iof = new(file03)
		}
		of.fileName = fileName

		if of.b, err = of.iof.readFile(of.fileName); err != nil {
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

func (me *officeFile) HideModule(moduleName string) (err error) {
	return me.vp.HideModule(moduleName)
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

	case "sm":
		if len(tokens) != 2 {
			fmt.Println(`sm <ModuleName> -- 保存模块代码为文件`)
			fmt.Printf("%d, %#v\r\n", len(tokens), tokens)
			return
		}

		moduleName := tokens[1]
		if index, err := of.vp.GetModuleIndex(moduleName); err != nil {
			fmt.Println(err)
		} else {
			if err := saveModule(index); err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%d 模块输出成功：%s\r\n", index, moduleName)
			}
		}

	case "smi":
		if len(tokens) != 2 {
			fmt.Println(`smi <ModuleIndex> -- 保存模块代码为文件`)
			fmt.Printf("%d, %#v\r\n", len(tokens), tokens)
			return
		}

		if index, err := strconv.Atoi(tokens[1]); err != nil {
			fmt.Println(err)
		} else {
			if err := saveModule(index); err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%d 模块输出成功：%s\r\n", index, of.vp.Module[index].Name)
			}
		}

	case "smAll":
		if len(tokens) != 1 {
			fmt.Println(`输入的命令不正确smAll -- 保存所有模块代码为文件`)
			fmt.Printf("%d, %#v\r\n", len(tokens), tokens)
			return
		}

		for i := range of.vp.Module {
			if err := saveModule(i); err != nil {
				fmt.Println(err)
				return
			} else {
				fmt.Printf("%d 模块输出成功：%s\r\n", i, of.vp.Module[i].Name)
			}
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

	case "hidem":
		if len(tokens) != 2 {
			fmt.Println(`输入的命令不正确hidem(ModuleName) -- 隐藏模块`)
			fmt.Printf("%d, %#v\r\n", len(tokens), tokens)
			return
		}

		if err := of.vp.HideModule(tokens[1]); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("隐藏模块%s成功\r\n", tokens[1])
			colorPrint.SetColor(colorPrint.White, colorPrint.Red)
			fmt.Println("退出前请保存文件：save")
		}

	case "unhidem":
		if len(tokens) < 2 {
			fmt.Println(`输入的命令不正确unhidem(ModuleName...) -- 取消多个隐藏模块`)
			fmt.Printf("%d, %#v\r\n", len(tokens), tokens)
			return
		}

		if err := of.vp.UnHideModule(tokens[1:]); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("取消隐藏模块%s成功\r\n", tokens[1:])
			colorPrint.SetColor(colorPrint.White, colorPrint.Red)
			fmt.Println("退出前请保存文件：save")
		}
	case "project":
		if err := of.vp.UnProtectProject(); err != nil {
			fmt.Println(err)
			return
		}

	case "save":
		var oldFileName string = of.fileName
		var saveFileName string
		if len(tokens) == 1 {
			saveFileName = of.fileName
		} else {
			saveFileName = tokens[1]
		}

		if err := of.iof.reWriteFile(oldFileName, saveFileName, of.b); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("保存成功。")
		}

	default:
		fmt.Println("Unrecognized lib command:", tokens)
	}
}

func saveModule(moduleIndex int) (err error) {
	var moduleName string = of.vp.Module[moduleIndex].Name

	var str string
	if str, err = of.vp.GetModuleCode(moduleName); err != nil {
		return
	}

	if of.vp.Module[moduleIndex].Type == vbaProject.CLASS_MODULE {
		moduleName += ".cls"
	} else {
		moduleName += ".bas"
	}

	if err = ioutil.WriteFile(moduleName, []byte(str), 0666); err != nil {
		return
	}

	return
}

func printCmd() {
	colorPrint.SetColor(colorPrint.Green, colorPrint.Black)

	fmt.Println(`
 Enter following commands to control:
 ls -- 查看模块列表
 project -- 破解vba密码
 show <ModuleName> -- 查看模块代码
 showi <ModuleIndex> -- 查看模块代码
 sm <ModuleName> -- 保存模块代码为文件
 smi <ModuleIndex> -- 保存模块代码为文件
 smAll -- 保存所有模块代码为文件
 hidem(ModuleName) -- 隐藏模块
 unhidem(ModuleName...) -- 取消多个隐藏模块
 save <[saveAsFileName默认是原文件名]>-- 保存文件
 e或者q -- 退出
 `)

	colorPrint.ReSetColor()
}
