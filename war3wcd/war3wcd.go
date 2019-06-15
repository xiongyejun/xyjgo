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
		fmt.Println("war3wcd <key> <mouseClick[0][1]> <PressCount>")
		return
	}

	var err error
	if PressCount, err = strconv.Atoi(os.Args[3]); err != nil {
		fmt.Println(err)
		return
	}

	key = os.Args[1][0:]
	if os.Args[2] == "1" {
		go handleMouse()
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
		if keve == 0 {
			for i := 0; i < PressCount; i++ {
				robotgo.KeyTap(key)
				robotgo.KeyTap("3")
			}
		}
	}
}
func handleMouse() {
	for {
		keve := robotgo.AddEvent(key)
		if keve == 0 {
			for i := 0; i < PressCount; i++ {
				robotgo.KeyTap(key)
				robotgo.MouseClick()
				robotgo.KeyTap("3")
			}
		}
	}
}
