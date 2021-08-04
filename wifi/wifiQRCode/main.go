package main

import (
	"fmt"
	"image"
	"os"

	"image/jpeg"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	qrcode "github.com/skip2/go-qrcode"
	"github.com/xiongyejun/xyjgo/wifi"
)

func uiInit() (err error) {
	myApp := app.New()

	myWin := myApp.NewWindow("wifi QRCode")

	entryAccount := widget.NewEntry()
	entryPsd := widget.NewEntry()
	img := canvas.NewImageFromImage(nil)

	funcShowImage := func() {
		var im image.Image
		if im, err = getQRCodeImage(entryAccount.Text, entryPsd.Text); err != nil {
			return
		}
		img.Image = im

		img.Refresh()
	}

	entryAccount.OnChanged = func(str string) {
		funcShowImage()
	}
	entryPsd.OnChanged = func(str string) {
		funcShowImage()
	}

	btn := widget.NewButton("get wifi", func() {
		if entryAccount.Text, err = wifi.GetSSID(); err != nil {
			return
		}

		entryAccount.Refresh()

		if entryPsd.Text, err = wifi.GetPsw(entryAccount.Text); err != nil {
			return
		}
		entryPsd.Refresh()

		funcShowImage()
	})

	lbtishi := widget.NewLabel("")
	btnSavePic := widget.NewButton("save pic", func() {
		savePic(img.Image)
		lbtishi.SetText("wifi qrcode name: qrcode.jpg")
	})

	content := widget.NewVBox(
		widget.NewHBox(
			widget.NewLabel("wifi  account:"),
			entryAccount,
		),

		widget.NewHBox(
			widget.NewLabel("wifi password:"),
			entryPsd,
		),

		btn,

		fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fyne.NewSize(300, 300)),
			img),

		btnSavePic,

		lbtishi,
	)

	myWin.SetContent(content)
	myWin.Resize(fyne.NewSize(200, 200))
	myWin.ShowAndRun()

	return
}

func main() {
	if err := uiInit(); err != nil {
		fmt.Println(err)
	}

}

func getQRCodeImage(account, psd string) (im image.Image, err error) {
	var strqrcode = wifi.QRCodeFormat(account, psd)

	var qr *qrcode.QRCode
	if qr, err = qrcode.New(strqrcode, qrcode.Low); err != nil {
		return
	}

	im = qr.Image(512)
	return
}

func savePic(im image.Image) (err error) {
	var f *os.File
	if f, err = os.Create("qrcode.jpg"); err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	if err = jpeg.Encode(f, im, &jpeg.Options{100}); err != nil {
		return
	}

	return
}
