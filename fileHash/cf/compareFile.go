// 对比2个文件夹里的文件
package main

import (
	"fmt"
	"os"

	"github.com/xiongyejun/xyjgo/colorPrint"
	"github.com/xiongyejun/xyjgo/fileHash"
)

func main() {
	print("Compare File 比较2个文件夹中的文件。\n")
	if len(os.Args) != 3 {
		fmt.Println("cf <Dir1> <Dir2>")
		return
	}

	var d1, d2 *fileHash.DirInfo = fileHash.New(os.Args[1]), fileHash.New(os.Args[2])
	var err error

	if err = d1.GetFilesInfo(); err != nil {
		fmt.Println(err)
		return
	}

	if err = d2.GetFilesInfo(); err != nil {
		fmt.Println(err)
		return
	}

	var ret []string
	if ret, err = Compare(d1, d2); err != nil {
		fmt.Println(err)
		return
	}

	if len(ret) == 0 {
		fmt.Println("2个文件夹的文件相同。")
		return
	}

	colorPrint.SetColor(colorPrint.Black, colorPrint.Yellow)
	print("<")
	print(d2.Path)
	print(">")
	colorPrint.ReSetColor()
	print("中的文件在")
	colorPrint.SetColor(colorPrint.Black, colorPrint.Yellow)
	print("<")
	print(d1.Path)
	print(">")
	colorPrint.ReSetColor()
	print("中不存在的文件：\n")

	colorPrint.SetColor(colorPrint.Green, colorPrint.Black)
	for i := range ret {
		fmt.Printf("%3d %s\n", i, ret[i])
	}
	colorPrint.ReSetColor()
}

// d2中的文件，在d1中不存在的
func Compare(d1, d2 *fileHash.DirInfo) (ret []string, err error) {
	ret = make([]string, 0)

	for i := range d2.FilesInfo {
		if _, ok := d1.MHashIndex[d2.FilesInfo[i].Hash]; !ok {
			ret = append(ret, d2.FilesInfo[i].FullName)
		}
	}

	return
}
