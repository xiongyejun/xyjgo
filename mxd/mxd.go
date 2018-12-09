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

	"github.com/xiongyejun/xyjgo/winAPI/user32/keyboard"

	"github.com/xiongyejun/xyjgo/colorPrint"
	"github.com/xiongyejun/xyjgo/pressKey"
)

type key struct {
	OnOff bool // 开关
	WVk   uint16
	Time  time.Duration // 间隔，毫秒
	Delay time.Duration // 按键后延迟时间，毫秒
	Note  string        // 备注说明
}

type skey struct {
	name string
	Keys []*key
	k    *pressKey.Key
	path string

	CountOn   int // 设置启动的个数
	bMove     bool
	bStop     bool
	moveValue uint16
	moveSleep time.Duration
}

const FILE_EXT string = ".keybd"

var s *skey

func init() {
	s = new(skey)
	s.k = pressKey.New()
	s.k.WindowText = "MapleStory"
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
	if s.name != "" && len(s.Keys) > 0 {
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

		if len(tokens) != 5 {
			fmt.Println("add <键, 间隔时间, 延时，备注说明> -- 增加1个按键")
			return
		}
		WVk := uint16([]byte(tokens[1])[0])
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
		Delay := time.Duration(tmp)

		s.Keys = append(s.Keys, &key{true, WVk, Time, Delay, tokens[4]})

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
			if index >= len(s.Keys) {
				fmt.Println("out of range")
				return
			}

			s.Keys = append(s.Keys[0:index], s.Keys[index+1:]...)
		}

	case "ls":
		if len(s.Keys) == 0 {
			fmt.Println("没有设置。")
			return
		}
		fmt.Printf("%2s %6s %4s %10s %10s %s\r\n", "No", "State", "key", "Time(ms)", "Delay(ms)", "Note")
		for i := range s.Keys {
			fmt.Printf("%2d %6s %4c %10d %10d %s\r\n", i, printOnOff(s.Keys[i].OnOff), s.Keys[i].WVk, s.Keys[i].Time, s.Keys[i].Delay, s.Keys[i].Note)
		}

	case "status":
		if s.name == "" {
			fmt.Println("还没有设置，先new1个或者read1个")
			return
		}

		if len(tokens) != 2 {
			fmt.Println("status <index> -- 开关")
			return
		}

		if index, err := strconv.Atoi(tokens[1]); err != nil {
			fmt.Println(err)
			return
		} else {
			if index >= len(s.Keys) {
				fmt.Println("out of range")
				return
			}

			s.Keys[index].OnOff = !s.Keys[index].OnOff
		}

	case "move":
		if s.name == "" {
			fmt.Println("move <moveSleep(秒)>-- 左右移动")
			return
		}

		if len(tokens) != 2 {
			fmt.Println("status <index> -- 开关")
			return
		}

		if i, err := strconv.Atoi(tokens[1]); err != nil {
			fmt.Println(err)
			return
		} else {
			if i > 0 {
				s.bMove = true
				s.moveSleep = time.Duration(i) * time.Second
			} else {
				fmt.Println("时间必须大于0")
			}
		}

	case "start":
		if len(s.Keys) == 0 {
			fmt.Println("没有设置。")
			return
		}

		s.bStop = false
		s.moveValue = keyboard.VK_RIGHT
		s.k = pressKey.New()
		s.k.WindowText = "MapleStory"
		s.CountOn = 0
		for i := range s.Keys {
			if s.Keys[i].OnOff {
				if s.Keys[i].WVk >= 'a' && s.Keys[i].WVk <= 'z' {
					s.Keys[i].WVk = s.Keys[i].WVk + 'A' - 'a'
				}
				s.k.Add(s.Keys[i].WVk, s.Keys[i].Time, s.Keys[i].Delay)
				s.CountOn++
			}
		}
		if s.bMove {
			s.k.Add(keyboard.VK_1, 1000, 1000) //攻击
			s.k.Add(s.moveValue, 1000, 1000)   // 移动

			fmt.Println("左右移动。")
			go s.move()
		}

		fmt.Println("3秒后开始按键。")
		time.Sleep(3 * time.Second)
		go s.k.Start()
	case "stop":
		s.k.Stop()
		s.bStop = true
	default:
		fmt.Println("不能存在的命令。")
	}
}

func printOnOff(b bool) string {
	if b {
		return "√"
	} else {
		return "×"
	}
}

func readKey(fileName string) (err error) {
	var b []byte
	if b, err = ioutil.ReadFile(fileName); err != nil {
		return
	}

	if err = json.Unmarshal(b, s); err != nil {
		return
	}

	return nil
}

func saveKey(fileName string) (err error) {
	var b []byte
	if b, err = json.MarshalIndent(s, "\r\n", "  "); err != nil {
		return
	}
	if err = ioutil.WriteFile(fileName, b, 0666); err != nil {
		return
	}
	return nil
}

// 左右移动，定时切换左右按键
func (me *skey) move() {
	index := me.CountOn + 1 // 攻击和移动增加的不再s.Keys里面

	for {
		time.Sleep(me.moveSleep)

		if me.bStop {
			return
		}

		if s.moveValue == keyboard.VK_LEFT {
			s.moveValue = keyboard.VK_RIGHT
		} else {
			s.moveValue = keyboard.VK_LEFT
		}

		if err := s.k.Change(index, s.moveValue); err != nil {
			s.k.Stop()
			s.bStop = true
			fmt.Println("Change err:", err)
		}
	}
}

func printCmd() {
	colorPrint.SetColor(colorPrint.Green, colorPrint.Black)

	fmt.Println(` Enter following commands to control:
 new <name> -- 新设置1个
 ls -- 查看当前设置的按键
 add <键, 间隔时间, 延时，备注说明> -- 增加1个按键
 del <index> -- 删除1个按键
 start -- 开始
 stop -- 结束
 status <index> -- 开关
 move <moveSleep(秒)>-- 左右移动
 read <name> -- 读取1个设置
`)
	colorPrint.ReSetColor()
}
