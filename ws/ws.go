// window style
// 设置窗口的style
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/xiongyejun/xyjgo/winAPI/user32"
)

type style struct {
	name  string
	text  string
	value int32
}

var styles []*style = []*style{
	&style{"WS_OVERLAPPED ", "说明", 0X00000000},
	//	&style{"WS_POPUP", "说明", 0X80000000},
	&style{"WS_CHILD", "说明", 0X40000000},
	&style{"WS_MINIMIZE ", "说明", 0X20000000},
	&style{"WS_VISIBLE", "说明", 0X10000000},
	&style{"WS_DISABLED ", "不允许点击", 0X08000000},
	&style{"WS_CLIPSIBLINGS ", "说明", 0X04000000},
	&style{"WS_CLIPCHILDREN ", "说明", 0X02000000},
	&style{"WS_MAXIMIZE ", "说明", 0X01000000},
	&style{"WS_CAPTION", "说明", 0X00C00000},
	&style{"WS_BORDER ", "说明", 0X00800000},
	&style{"WS_DLGFRAME ", "说明", 0X00400000},
	&style{"WS_VSCROLL", "说明", 0X00200000},
	&style{"WS_HSCROLL", "说明", 0X00100000},
	&style{"WS_SYSMENU", "说明", 0X00080000},
	&style{"WS_THICKFRAME ", "说明", 0X00040000},
	&style{"WS_GROUP", "说明", 0X00020000},
	&style{"WS_TABSTOP", "说明", 0X00010000},
	&style{"WS_MINIMIZEBOX", "说明", 0X00020000},
	&style{"WS_MAXIMIZEBOX", "说明", 0X00010000},
	&style{"WS_TILED", "说明", 0X00000000},
	&style{"WS_ICONIC ", "说明", 0X20000000},
	&style{"WS_SIZEBOX", "说明", 0X00040000},
	&style{"WS_OVERLAPPEDWINDOW ", "说明", 0X00000000 | 0X00C00000 | 0X00080000 | 0X00040000 | 0X00020000 | 0X00010000},
	//	&style{"WS_POPUPWINDOW", "说明", 0X80000000 | 0X00800000 | 0X00080000},
	&style{"WS_CHILDWINDOW", "说明", 0X40000000},
}

func main() {
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

func handleCommands(tokens []string) {
	switch tokens[0] {
	case "add":
		if err := add_del(tokens, "add <hwnd> <style index>-- 为hwnd窗口添加1个style", addStyle); err != nil {
			fmt.Println(err)
		}
	case "del":
		if err := add_del(tokens, "del <hwnd> <style index> -- 为hwnd窗口删除1个style", delStyle); err != nil {
			fmt.Println(err)
		}
	case "ls":
		printStyles()
	default:
		fmt.Println("Unrecognized lib command:", tokens)
	}
}

func printCmd() {
	fmt.Println(`
 Enter following commands to control:
 add <hwnd> <style index>-- 为hwnd窗口添加1个style
 del <hwnd> <style index> -- 为hwnd窗口删除1个style
 ls -- 查看styles
 e或者q -- 退出 
 `)
}

// 抽象添加和删除函数，共用1个函数完成
func add_del(tokens []string, strCmd string, f func(uint32, int32) error) (err error) {
	if len(tokens) != 3 {
		return errors.New(strCmd)
	} else {
		var hwnd int64
		if hwnd, err = strconv.ParseInt(tokens[1], 0, 32); err != nil {
			return
		} else {
			var styleindex int
			if styleindex, err = strconv.Atoi(tokens[2]); err != nil {
				return
			} else {
				if err = f(uint32(hwnd), styles[styleindex].value); err != nil {
					return
				}
			}
		}
	}
	return
}

// 添加1个样式
func addStyle(hwnd uint32, style int32) (err error) {
	stylesrc := user32.GetWindowLong(hwnd, user32.GWL_STYLE)
	stylesrc |= style
	user32.SetWindowLong(hwnd, user32.GWL_STYLE, stylesrc)

	return
}

// 删除1个样式
func delStyle(hwnd uint32, style int32) (err error) {
	fmt.Printf("hwnd=%x,style=%x\r\n", hwnd, style)

	stylesrc := user32.GetWindowLong(hwnd, user32.GWL_STYLE)
	stylesrc &= (^style)
	user32.SetWindowLong(hwnd, user32.GWL_STYLE, stylesrc)

	return
}

// 输出styles
func printStyles() {
	for i := range styles {
		fmt.Printf("%2d  %-25s %s\r\n", i, styles[i].name, styles[i].text)
	}
}

func getDesHwnd(wins []string) (hwnd uint32, err error) {
	hwnd = user32.FindWindow("", wins[0])

	if hwnd == 0 {
		return 0, errors.New("[" + wins[0] + "] 窗口句柄获取出错。")
	}
	for i := 1; i < len(wins); i++ {
		hwnd = user32.FindWindowEx(hwnd, 0, "", wins[i])
		if hwnd == 0 {
			return 0, errors.New("[" + wins[i] + "] 窗口句柄获取出错。")
		}
	}

	return
}
