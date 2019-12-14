// Optical Character Recognition，光学字符识别
package main

import (
	"fmt"
	"image"
)

type IOCR interface {
	OCR(picPath string) (ret string, err error)
}

func main() {

}
