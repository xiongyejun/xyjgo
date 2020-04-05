package main

import (
	"fmt"
	"os"

	"github.com/disintegration/imaging"
)

func main() {
	if len(os.Args) != 3 {
		printHelp()
		return
	}
	fmt.Println(os.Args[2])
	f, err := os.OpenFile(os.Args[2], os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	im, err := imaging.Decode(f)
	if err != nil {
		fmt.Println(err, 1)
		return
	}

	nrgba := imaging.Resize(im, im.Bounds().Dx(), int(float64(im.Bounds().Dx())/2.35), imaging.Box)

	err = imaging.Save(nrgba, "resize_"+os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("OK")
}

func printHelp() {
	fmt.Println(`
  pic dyh <path>    resize微信订阅号用的封面图片2.35:1
	`)
}
