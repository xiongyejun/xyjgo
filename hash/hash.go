// hash <hashType> <src> <[check hash]>
// 计算文本\文件的hash值
// <src> 先判断是否存在这个名称的文件，如果存在就计算文件的hash
// <[check hash]> 如果输入了第4个参数，最后进行hash值的判断

package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"io/ioutil"
	"os"
	"strings"

	"github.com/xiongyejun/xyjgo/colorPrint"
)

func main() {
	fmt.Println("hash <hashType> <src> <[check hash]>")
	if len(os.Args) < 3 {
		return
	}

	var strHash string = os.Args[1]
	strHash = strings.ToLower(strHash)

	var h hash.Hash
	if strHash == "sha1" {
		h = sha1.New()
	} else if strHash == "md5" {
		h = md5.New()
	} else if strHash == "sha256" {
		h = sha256.New()
	} else {
		fmt.Printf("未知的hash:%s\r\n", strHash)
		fmt.Println("请输入以下hash：\r\nsha1\r\nmd5\r\nsha256")
		return
	}

	var src string = os.Args[2]
	var bSrc []byte

	tmp, err := isWhat(src)

	if err != nil {
		fmt.Println(err)
		return
	}

	if tmp == OTHER || tmp == ISDIR {
		bSrc = []byte(src)
		print("文本")
	} else {
		// 如果文件过大，容易卡死，应该分开读取
		if bSrc, err = ioutil.ReadFile(src); err != nil {
			fmt.Println(err)
		} else {
			print("文件")
		}
	}

	if _, err := h.Write(bSrc); err != nil {
		fmt.Println(err)
	} else {
		b := h.Sum(nil)
		fmt.Printf(" [%s] hash = %x\r\n", src, b)

		if len(os.Args) == 4 {
			if strings.ToLower(os.Args[3]) == fmt.Sprintf("%x", b) {
				fmt.Println("与输入的hash值一致。")
			} else {
				colorPrint.SetColor(colorPrint.White, colorPrint.Red)
				fmt.Println("!!!与输入的hash值不一致。")
				colorPrint.ReSetColor()
			}
		}
	}
}

const (
	ISDIR = iota
	ISFILE
	OTHER
)

func isWhat(path string) (int, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return ISDIR, nil
		} else {
			return ISFILE, nil
		}
	}
	if os.IsNotExist(err) {
		return OTHER, nil
	}

	return OTHER, err
}
