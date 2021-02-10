package main

import (
	"fmt"
	"os/exec"

	"github.com/xiongyejun/xyjgo/wifi"
)

func main() {
	if ssid, err := wifi.GetSSID(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ssid)

		if psd, err := wifi.GetPsw(ssid); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(psd)
		}
	}

	exec.Command("cmd", "/C", `CHCP 936`)
}
