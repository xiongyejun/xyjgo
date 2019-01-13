package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"os"
	"strings"

	"github.com/golang/image/bmp"
)

var strDir string
var bFile bool

func init() {
	str, _ := os.Getwd() // 获得cmd命令行cd的路径
	tmpstrDir := flag.String("d", str+string(os.PathSeparator), "搜索当前文件夹下所有.bmp")
	tmpbFile := flag.Bool("f", false, "按指定的文件")
	flag.Parse()

	strDir = *tmpstrDir
	bFile = *tmpbFile

	flag.PrintDefaults()

}

func main() {
	if bFile {
		if len(os.Args) != 3 {
			fmt.Println("bmpTOjpg <-f> <srcFile>")
			return
		}
		if err := bmpTOjpeg(os.Args[2]); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println(os.Args[2] + " 转换成功")
		}
	} else {
		scanDir(strDir)
	}

	fmt.Println("ok")
}

func scanDir(dirName string) {
	entrys, err := ioutil.ReadDir(dirName)
	if err != nil {
		return
	}

	for _, entry := range entrys {
		if !entry.IsDir() {
			if strings.HasSuffix(entry.Name(), ".bmp") {
				if err := bmpTOjpeg(dirName + entry.Name()); err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(entry.Name() + " 转换成功")
				}
			}

		}
	}
}

func bmpTOjpeg(srcFile string) error {
	f, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer f.Close()
	img, err := bmp.Decode(f)
	if err != nil {
		return err
	}

	index := strings.LastIndex(srcFile, ".")
	desFile := srcFile[:index] + ".jpg"

	fjpg, err := os.Create(desFile)
	if err != nil {
		return err
	}
	defer fjpg.Close()

	err = jpeg.Encode(fjpg, img, nil)
	if err != nil {
		return err
	}
	return nil
}
