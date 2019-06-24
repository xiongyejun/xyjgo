package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-vgo/robotgo"
)

var key string = ""
var PressCount int = 0

func main() {
	if len(os.Args) != 4 {
		fmt.Println("war3 <handle key> <PressCount> <[Mouse(Left=0,Right=1,其他不按)]>")
		return
	}

	var err error
	if PressCount, err = strconv.Atoi(os.Args[2]); err != nil {
		fmt.Println(err)
		return
	}

	key = os.Args[1][0:]
	fmt.Println(key)

	if os.Args[3] == "1" {
		go handleMouse("right")
	} else if os.Args[3] == "0" {
		go handleMouse("left")
	} else {
		go handleNoMouse()
	}

	for {
		time.Sleep(10 * time.Second)
	}
}
func handleNoMouse() {
	for {
		keve := robotgo.AddEvent(key)
		if keve {
			for i := 0; i < PressCount; i++ {
				robotgo.KeyTap(key)
				time.Sleep(time.Second / 10)
				fmt.Println(key)
			}
		}
	}
}
func handleMouse(strMouse string) {
	for {
		keve := robotgo.AddEvent(key)
		if keve {
			for i := 0; i < PressCount; i++ {
				robotgo.MouseClick(strMouse)
				time.Sleep(time.Second / 10)
			}
		}
	}
}
