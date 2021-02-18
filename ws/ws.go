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
	&style{"WS_OVERLAPPED ", "产生一个层叠的窗口，一个层叠的窗口有一个标题栏和一个边 框。", 0X00000000},
	//	&style{"WS_POPUP", "创建一个弹出式窗口，不能与WS_CHILD风格一起使用", 0X80000000},
	&style{"WS_CHILD", "说明窗口为子窗口，不能应用于弹出式窗口风格(WS_POPUP)。", 0X40000000},
	&style{"WS_CHILDWINDOW", "同WS_CHILD", 0X40000000},
	&style{"WS_MINIMIZE ", "创建一个初始状态为最小化的窗口。仅与WS_OVERLAPPED风格一起使用", 0X20000000},
	&style{"WS_MAXIMIZE ", "创建一个最大化的窗口", 0X01000000},
	&style{"WS_VISIBLE", "创建一个最初可见的窗口", 0X10000000},
	&style{"WS_DISABLED ", "不允许点击", 0X08000000},
	&style{"WS_CLIPSIBLINGS ", "S 剪裁相关的子窗口，这意味着，当一个特定的子窗口接收到重绘消息时，WS_CLIPSIBLINGS风格将在子窗口要重画的区域中去掉与其它子窗口重叠的部分。（如果没有指定WS_CLIPSIBLINGS风格，并且子窗口有重叠，当你在一个子窗口的客户区绘图时，它可能会画在相邻的子窗口的客户区中。）只与WS_CHILD风格一起使用。", 0X04000000},
	&style{"WS_CLIPCHILDREN ", "绘制父窗口时，不绘制子窗口的裁剪区域。使用在建立父窗口时", 0X02000000},
	&style{"WS_CAPTION", "创建一个有标题栏的窗口", 0X00C00000},
	&style{"WS_BORDER ", "有边框窗口", 0X00800000},
	&style{"WS_DLGFRAME ", "创建一个窗口，具有双重边界，但是没有标题条", 0X00400000},
	&style{"WS_HSCROLL", "创建一个具有水平滚动条的窗口", 0X00100000},
	&style{"WS_VSCROLL", "创建一个具有垂直滚动条的窗口", 0X00200000},
	&style{"WS_SYSMENU", "创建一个在标题栏上带有系统菜单的窗口，要和WS_CAPTION类 型一起使用", 0X00080000},
	&style{"WS_THICKFRAME ", "创建一个具有可调边框的窗口。", 0X00040000},
	&style{"WS_GROUP", "指定一组控件中的第一个，用户可以用箭头键在这组控件中移动。在第一个控件后面把WS_GROUP风格设置为FALSE的控件都属于这一组。下一个具有WS_GROUP风格的控件将开始下一组（这意味着一个组在下一组的开始处结束）", 0X00020000},
	&style{"WS_TABSTOP", "指定了一些控件中的一个，用户可以通过TAB键来移过它。TAB键使用户移动到下一个用WS_TABSTOP风格定义的控件", 0X00010000},
	&style{"WS_MINIMIZEBOX", "创建一个具有最小化按钮的窗口，必须同时设定WS_ SYSMENU类型", 0X00020000},
	&style{"WS_MAXIMIZEBOX", "创建一个具有最大化按钮的窗口，必须同时设定WS_ SYSMENU类型", 0X00010000},
	&style{"WS_TILED", "产生一个层叠的窗口。一个层叠的窗口有一个标题和一个边框。与WS_OVERLAPPED风格相同", 0X00000000},
	&style{"WS_ICONIC ", "创建一个初始状态为最小化状态的窗口。与WS_MINIMIZE风格相同。", 0X20000000},
	&style{"WS_SIZEBOX", "创建一个可调边框的窗口，与WS_THICKFRAME风格相同", 0X00040000},
	&style{"WS_OVERLAPPEDWINDOW ", "创建一个具有WS_OVERLAPPED,WS_CAPTION,WS_SYSMENU,WS_THICKFRAME,WS_MINIMIZEBOX和WS_MAXIMIZEBOX风格的重叠式窗口。 ", 0X00000000 | 0X00C00000 | 0X00080000 | 0X00040000 | 0X00020000 | 0X00010000},
	//	&style{"WS_POPUPWINDOW", "创建一个具有WS_BORDER，WS_POPUP和WS_SYSMENU风格的弹出窗口。为了使控制菜单可见，必须与WS_POPUPWINDOW一起使用WS_CAPTION风格。 ", 0X80000000 | 0X00800000 | 0X00080000},
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

	case "fw":
		if len(tokens) != 2 {
			printCmd()
			return
		}
		fmt.Printf("窗口[%s] hwnd = [%d]\n", tokens[1], findWindow(tokens[1]))
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
 fw <sname>  -- 查找窗口hwnd
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

func findWindow(sname string) uint32 {
	return user32.FindWindow("", sname)

}
