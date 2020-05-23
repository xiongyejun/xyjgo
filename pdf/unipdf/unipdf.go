package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TODO 记得修改这个
/*
github.com\unidoc\unipdf\common\license\key.go
func (k *LicenseKey) IsLicensed() bool {
	// return k.Tier != LicenseTierUnlicensed
	return true
}
*/

func main() {
	if len(os.Args) == 1 {
		printHelp()
		return
	}
	var err error
	switch os.Args[1] {
	case "i":
		// i <imagesFolder> <outputPath> -- images to pdf 图片创建pdf
		if len(os.Args) != 4 {
			printHelp()
		}
		if err = imagesToPdf(os.Args[2], os.Args[3]); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("images to pdf OK.")

	case "s":
		// s <inputPath> <outputPath> <pages> -- 拆分pdf
		if len(os.Args) != 5 {
			printHelp()
		}
		if err = splitPdf(os.Args[2], os.Args[3], os.Args[4]); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Split OK.")

	case "c":
		// c <inputPath> <outputPath> -- Compress and optimize PDF 压缩pdf
		if len(os.Args) != 4 {
			printHelp()
		}
		if err = compress(os.Args[2], os.Args[3]); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Compress OK.")

	case "e":
		// e <inputPath> <outputPath> -- Extract images 提取图片
		if len(os.Args) != 4 {
			printHelp()
		}
		if err = extractImages(os.Args[2], os.Args[3]); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Extract OK.")

	case "m":
		// m <outputPath> <inputPaths...> -- pdf merge 合并pdf
		if len(os.Args) < 4 {
			printHelp()
		}
		if err = merge(os.Args[2], os.Args[3:]); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Merge OK.")

	case "r":
		// r <inputPaths> <outputPath> <degrees> <pages> -- rotate旋转pdf，degrees角度(90度的整数)，page base 1
		if len(os.Args) != 6 {
			printHelp()
		}
		if err = rotate(os.Args[2], os.Args[3], os.Args[4], os.Args[5]); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Rotate OK.")
	default:
		printHelp()
	}
}

// <pages> 逗号分隔，支持1-10形式
func getPages(pages string) (page []int, err error) {
	arr := strings.Split(pages, ",")

	var tmp int
	for i := range arr {
		tmparr := strings.Split(arr[i], "-")
		if len(tmparr) == 1 {
			if tmp, err = strconv.Atoi(tmparr[0]); err != nil {
				return
			}
			page = append(page, tmp)
		} else if len(tmparr) == 2 {
			var fromP, toP int
			if fromP, err = strconv.Atoi(tmparr[0]); err != nil {
				return
			}
			if toP, err = strconv.Atoi(tmparr[1]); err != nil {
				return
			}

			for j := fromP; j <= toP; j++ {
				page = append(page, j)
			}
		} else {
			err = errors.New("不能多于1个[-]")
			return
		}

	}

	return
}

func printHelp() {
	fmt.Println(`
 i <imagesFolder> <outputPath> -- images to pdf 图片创建pdf
 s <inputPath> <outputPath> <pages> -- 拆分pdf
 c <inputPath> <outputPath> -- Compress and optimize PDF 压缩pdf
 e <inputPath> <outputPath> -- Extract images 提取图片
 m <outputPath> <inputPaths...> -- pdf merge 合并pdf
 r <inputPaths> <outputPath> <degrees> <pages> -- rotate旋转pdf，degrees角度(90度的整数)，page base 1

 <pages> 逗号分隔，支持1-10形式
	`)
}
