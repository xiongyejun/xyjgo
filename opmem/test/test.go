package main

import (
	"fmt"

	"github.com/xiongyejun/xyjgo/opmem"
)

func main() {
	ret, err := opmem.NewByWindow(`C:\Users\Administrator\Documents\08-go\src\github.com\xiongyejun\xyjgo\opmem\exe32\exe32.exe`)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Process)

	retv, err := ret.Scan()
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range retv {
		b, err := ret.Read(retv[i])
		if err != nil {
			fmt.Printf("%x, %s\n", retv[i].BaseAddress, err.Error())
		} else {
			fmt.Printf("%x, size=%x, value= % x\n", retv[i].BaseAddress, retv[i].RegionSize, b[:10])
		}
	}

}
