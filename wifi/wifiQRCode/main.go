package main

import (
	"fmt"
	"image"
	"os"

	_ "embed"
	"image/jpeg"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	qrcode "github.com/skip2/go-qrcode"
	"github.com/xiongyejun/xyjgo/wifi"
)

//go:embed 扫码_搜索联合传播样式-标准色版.png
var bxyjvba []byte

func uiInit() (err error) {
	myApp := app.New()

	myWin := myApp.NewWindow("wifi QRCode")

	lbtishi := widget.NewLabel("")
	// 帐号
	entryAccount := widget.NewEntry()
	// 密码
	entryPsd := widget.NewEntry()
	// 显示二维码
	img := canvas.NewImageFromImage(nil)

	// img控件显示二维码
	funcShowImage := func() {
		var im image.Image
		if im, err = getQRCodeImage(entryAccount.Text, entryPsd.Text); err != nil {
			return
		}
		img.Image = im

		img.Refresh()
	}
	// 文本改变了就更新img控件
	entryAccount.OnChanged = func(str string) {
		funcShowImage()
	}
	entryPsd.OnChanged = func(str string) {
		funcShowImage()
	}

	// 获取活动的wifi账号密码
	btn := widget.NewButton("get wifi", func() {
		if entryAccount.Text, err = wifi.GetSSID(); err != nil {
			lbtishi.SetText(err.Error())
			return
		}

		entryAccount.Refresh()

		if entryPsd.Text, err = wifi.GetPsw(entryAccount.Text); err != nil {
			lbtishi.SetText(err.Error())
			return
		}
		entryPsd.Refresh()

		funcShowImage()
	})

	// 保存图片后提示信息

	// 保存图片按钮
	btnSavePic := widget.NewButton("save pic", func() {
		if img.Image == nil {
			lbtishi.SetText("no image!")
			return
		}
		if err = savePic(img.Image); err != nil {
			lbtishi.SetText(err.Error())
			return
		}
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

		widget.NewLabel("contact me: qq648555205"),
		fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fyne.NewSize(300, 100)),
			canvas.NewImageFromResource(xyjvba)),
	)

	myWin.SetContent(content)
	myWin.Resize(fyne.NewSize(200, 200))
	myWin.ShowAndRun()

	return
}

func main() {
	xyjvba.name = "xyjvba"
	xyjvba.b = bxyjvba

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
		return
	}
	defer f.Close()

	if err = jpeg.Encode(f, im, &jpeg.Options{100}); err != nil {
		return
	}

	return
}

var xyjvba gzh = gzh{}

type gzh struct {
	name string
	b    []byte
}

func (me gzh) Name() string {
	return me.name
}

func (me gzh) Content() []byte {
	return me.b
}
