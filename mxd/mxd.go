package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/xiongyejun/xyjgo/colorPrint"
	"github.com/xiongyejun/xyjgo/pressKey"
)

type skey struct {
	name string
	k    *pressKey.Key
	path string
}

const FILE_EXT string = ".keybd"

var s *skey

func init() {
	s = new(skey)
	s.k = pressKey.New()
	s.path = os.Getenv("GOPATH") + `\src\github.com\xiongyejun\xyjgo\mxd\`
}

func main() {

	printCmd()
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter Cmd->")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)
		if line == "q" || line == "e" {
			s.k.Stop()
			break
		}
		tokens := strings.Split(line, " ")
		printCmd()
		handleCommands(tokens)
	}

	// 退出的时候保存当前的设置
	if s.name != "" && len(s.k.Keys) > 0 {
		if err := saveKey(s.path + s.name + FILE_EXT); err != nil {
			fmt.Println(err)
		}
	}
}

func handleCommands(tokens []string) {
	switch tokens[0] {
	case "new":
		if len(tokens) != 2 {
			fmt.Println("new <name> -- 新设置1个")
			return
		}
		s.name = tokens[1]

	case "read":
		if len(tokens) != 2 {
			fmt.Println("read <name> -- 读取1个设置")
			return
		}

		if _, err := os.Stat(s.path + tokens[1] + FILE_EXT); err != nil {
			fmt.Println("不存在的文件:", tokens[1])
			return
		}
		s.name = tokens[1]
		// 读取设置-json保存
		if err := readKey(s.path + s.name + FILE_EXT); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("读取成功。")
		}

	case "add":
		if s.name == "" {
			fmt.Println("还没有设置，先new1个或者read1个")
			return
		}

		if len(tokens) != 4 {
			fmt.Println("add <键, 间隔时间, 延时> -- 增加1个按键")
			return
		}
		WVk := int([]byte(tokens[1])[0])
		var tmp int
		var err error
		if tmp, err = strconv.Atoi(tokens[2]); err != nil {
			fmt.Println(err)
			return
		}
		Time := time.Duration(tmp)

		if tmp, err = strconv.Atoi(tokens[3]); err != nil {
			fmt.Println(err)
			return
		}
		Delay := uint(tmp)

		if err = s.k.Add(WVk, Time, Delay); err != nil {
			fmt.Println(err)
			return
		}

	case "del":
		if s.name == "" {
			fmt.Println("还没有设置，先new1个或者read1个")
			return
		}

		if len(tokens) != 2 {
			fmt.Println("del <index> -- 删除1个按键")
			return
		}

		if index, err := strconv.Atoi(tokens[1]); err != nil {
			fmt.Println(err)
			return
		} else {
			if err = s.k.Remove(index); err != nil {
				fmt.Println(err)
				return
			}
		}

	case "ls":
		if len(s.k.Keys) == 0 {
			fmt.Println("没有设置。")
			return
		}
		fmt.Printf("%2s %4s %10s %s\r\n", "No", "key", "Time(ms)", "Delay(ms)")
		for i := range s.k.Keys {
			fmt.Printf("%2d %4c %10d %d\r\n", i, s.k.Keys[i].WVk, s.k.Keys[i].Time, s.k.Keys[i].Delay)
		}

	case "start":
		fmt.Println("3秒后开始按键。")
		time.Sleep(3 * time.Second)
		go s.k.Start()
	case "stop":
		s.k.Stop()
	default:
		fmt.Println("不能存在的命令。")
	}
}

func readKey(fileName string) (err error) {
	var b []byte
	if b, err = ioutil.ReadFile(fileName); err != nil {
		return
	}
	s.k = pressKey.New()
	if err = json.Unmarshal(b, s.k); err != nil {
		return
	}
	return nil
}
func saveKey(fileName string) (err error) {
	var b []byte
	if b, err = json.Marshal(s.k); err != nil {
		return
	}
	if err = ioutil.WriteFile(fileName, b, 0666); err != nil {
		return
	}
	return nil
}

func printCmd() {
	colorPrint.SetColor(colorPrint.Green, colorPrint.Black)

	fmt.Println(` Enter following commands to control:
 new <name> -- 新设置1个
 ls -- 查看当前设置的按键
 add <键, 间隔时间, 延时> -- 增加1个按键
 del <index> -- 删除1个按键
 start -- 开始
 stop -- 结束
 read <name> -- 读取1个设置
`)
	colorPrint.ReSetColor()
}
