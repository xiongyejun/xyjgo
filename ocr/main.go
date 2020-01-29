// Optical Character Recognition，光学字符识别
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/xiongyejun/xyjgo/ocr/baidu"
)

type IOCR interface {
	OCR(picPath string) (ret string, err error)
}

var iocr IOCR
var strDir string
var bFile bool

func init() {
	str, _ := os.Getwd() // 获得cmd命令行cd的路径
	tmpstrDir := flag.String("d", str+string(os.PathSeparator), "搜索当前文件夹下所有图片")
	tmpbFile := flag.Bool("f", false, "按指定的文件")
	flag.Parse()

	strDir = *tmpstrDir + string(os.PathSeparator)
	bFile = *tmpbFile

	flag.PrintDefaults()

}

func main() {
	var err error
	if iocr, err = baidu.New(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(strDir)

	if bFile {
		if len(os.Args) != 3 {
			fmt.Println("ocr <-f> <srcFile>")
			return
		}
		ocrToFile(os.Args[2])

	} else {
		if len(os.Args) != 3 {
			fmt.Println("ocr <-d> <dir> 识别文件夹里的所有文件，并保存txt")
			return
		}

		if files, err := scanDir(strDir); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(files)

			if len(files) == 0 {
				fmt.Println("没有图片。")
				return
			}
			var ch chan int = make(chan int, len(files))
			for i := range files {
				go func(fileName string) {
					ocrToFile(fileName)
					ch <- i
				}(files[i])
			}
			for i := 0; i < len(files); i++ {
				<-ch
			}
		}
	}
}

func ocrToFile(picName string) {
	if ret, err := iocr.OCR(picName); err != nil {
		fmt.Println(err)
	} else {
		if err := ioutil.WriteFile(picName+".txt", []byte(ret), 0666); err != nil {
			fmt.Println(err)
		}
	}
}

func scanDir(dirName string) (ret []string, err error) {
	var entrys []os.FileInfo
	if entrys, err = ioutil.ReadDir(dirName); err != nil {
		return
	}

	for _, entry := range entrys {
		if !entry.IsDir() {
			switch filepath.Ext(entry.Name()) {
			case ".jpg", ".jpeg", ".png":
				ret = append(ret, dirName+entry.Name())
			}
		}
	}

	return
}
