package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Comdex/imgo"
	"github.com/disintegration/imaging"
)

func main() {
	var err error
	if len(os.Args) == 1 {
		printHelp()
		return
	}

	switch os.Args[1] {
	case "dyh":
		if err = Resize(os.Args[2]); err != nil {
			fmt.Println(err)
			return
		}
	case "bz":
		threshold := 127
		if len(os.Args) == 4 {
			if threshold, err = strconv.Atoi(os.Args[3]); err != nil {
				fmt.Println(err)
				return
			}

			if err = Binaryzation(os.Args[2], threshold); err != nil {
				fmt.Println(err)
				return
			}
		}

	default:
		printHelp()
	}

	fmt.Println("OK")
}

func Resize(picname string) (err error) {
	f, err := os.OpenFile(picname, os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	im, err := imaging.Decode(f)
	if err != nil {
		return
	}

	nrgba := imaging.Resize(im, im.Bounds().Dx(), int(float64(im.Bounds().Dx())/2.35), imaging.Box)

	return imaging.Save(nrgba, "resize_"+picname)
}

// 二值化
func Binaryzation(picname string, threshold int) (err error) {
	img := imgo.MustRead(picname)
	img = imgo.Binaryzation(img, threshold)
	return imgo.SaveAsJPEG("bz_"+picname, img, 100)
}

func printHelp() {
	fmt.Println(`
  pic dyh <path>    resize微信订阅号用的封面图片2.35:1
  pic bz <path> <[threshold=127]>    Binaryzation二值化
	`)
}
