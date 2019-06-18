package main

import (
	//	"fmt"
	"math"

	"github.com/lxn/walk"

	qrcode "github.com/skip2/go-qrcode"
)

// 字符分段处理
type regionStr struct {
	start int
	end   int
}

type imageInfo struct {
	str        string
	b          []byte
	region_Str []regionStr

	next int
}

const MAX_B float64 = 1600 // 太大了容易扫不出来

// 1个图片最多2048个字节
func getImage(str string) (images []walk.Image) {
	image_Info := new(imageInfo)
	image_Info.b = []byte(str)
	// 需要的图片数量
	nImage := math.Ceil(float64(len(image_Info.b)) / MAX_B)
	// 每张图片的字节数，平分一下
	nByte := int(math.Ceil(float64(len(image_Info.b)) / nImage))

	image_Info.region_Str = make([]regionStr, int(nImage))

	// 字符分段处理，需要注意utf编码的不一定字节情况
	for i, _ := range str {
		if i >= (nByte + image_Info.next*nByte) {
			image_Info.region_Str[image_Info.next].end = i

			image_Info.next++
			if image_Info.next < int(nImage) {
				image_Info.region_Str[image_Info.next].start = i
			}
		}
	}
	if image_Info.next < int(nImage) {
		image_Info.region_Str[image_Info.next].end = len(image_Info.b)
	}

	return image_Info.getQRCode()
}

func (me *imageInfo) getQRCode() (images []walk.Image) {
	for i := range me.region_Str {
		str := string(me.b[me.region_Str[i].start:me.region_Str[i].end])

		qr, err := qrcode.New(str, qrcode.Low)
		if err != nil {
			return nil
		}
		im, err := walk.NewBitmapFromImage(qr.Image(512))
		if err != nil {
			return nil
		}
		images = append(images, im)
	}
	return
}
