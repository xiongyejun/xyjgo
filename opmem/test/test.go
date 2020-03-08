package main

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/xiongyejun/xyjgo/opmem"
)

func main() {
	ret, err := opmem.NewByProcessName(`exe32.exe`)
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
			for j := 0; j < len(b); j += 4 {
				if bytes.Compare(b[j:j+4], []byte{0, 0, 0, 0}) != 0 {
					fmt.Printf("%x = %d\n", retv[i].BaseAddress+int32(j), binary.LittleEndian.Uint32(b[j:j+4]))
				}

			}
		}
	}

}
