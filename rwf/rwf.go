// ReWriteFile 改写文件
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("rwf <fileName> <Pos(16进制)> <String>")
		return
	}

	fileName := os.Args[1]
	var iPos int64
	var err error
	if iPos, err = strconv.ParseInt(os.Args[2], 16, 64); err != nil {
		fmt.Println(err)
		return
	}

	var f *os.File
	if f, err = os.OpenFile(fileName, os.O_RDONLY|os.O_WRONLY, 0666); err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	if _, err = f.Seek(iPos, 0); err != nil {
		fmt.Println(err)
		return
	}

	if _, err = f.Write([]byte(os.Args[3])); err != nil {
		fmt.Println(err)
		return
	}

}
